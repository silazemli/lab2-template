package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/silazemli/lab2-template/internal/services/gateway/clients"
	"github.com/silazemli/lab2-template/internal/services/payment"
	"github.com/silazemli/lab2-template/internal/services/reservation"
)

type server struct {
	srv         echo.Echo
	cfg         Config
	reservation clients.ReservationClient
	payment     clients.PaymentClient
	loyalty     clients.LoyaltyClient
}

func NewServer() server {
	srv := server{}
	srv.srv = *echo.New()
	srv.cfg = *NewConfig()

	client := &http.Client{
		Transport: &http.Transport{MaxConnsPerHost: 100},
		Timeout:   5 * time.Second,
	}
	srv.loyalty = *clients.NewLoyaltyClient(client, srv.cfg.LoyaltyService)
	srv.payment = *clients.NewPaymentClient(client, srv.cfg.PaymentService)
	srv.reservation = *clients.NewReservationClient(client, srv.cfg.ReservationService)

	api := srv.srv.Group("/api/v1")
	api.GET("/hotels", srv.GetAllHotels)                               // +
	api.GET("/me", srv.GetUser)                                        // +
	api.GET("/loyalty", srv.GetStatus)                                 // +
	api.GET("/reservations", srv.GetAllReservations)                   // +
	api.GET("/reservations/:reservationUid", srv.GetReservation)       // +
	api.POST("/reservations", srv.MakeReservation)                     // +
	api.DELETE("/reservations/:reservationUid", srv.CancelReservation) // +

	srv.srv.GET("/manage/health", srv.HealthCheck)

	return srv
}

func (srv *server) Start() error {
	err := srv.srv.Start(":8080")
	if err != nil {
		return err
	}
	return nil
}

func (srv *server) GetUser(ctx echo.Context) error {
	username := ctx.Request().Header.Get("X-User-Name")
	user, err := srv.loyalty.GetUser(username)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusOK, user)
}

func (srv *server) GetAllHotels(ctx echo.Context) error {
	hotels, err := srv.reservation.GetAllHotels()
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusOK, hotels)
}

func (srv *server) GetAllReservations(ctx echo.Context) error {
	username := ctx.Request().Header.Get("X-User-Name")
	reservations, err := srv.reservation.GetReservations(username)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusOK, reservations)
}

func (srv *server) GetReservation(ctx echo.Context) error {
	username := ctx.Request().Header.Get("X-User-Name")
	reservationUID := ctx.Param("reservationUid")
	theReservation, err := srv.reservation.GetReservation(reservationUID)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	if username != theReservation.Username {
		return ctx.JSON(http.StatusForbidden, echo.Map{"error": err})
	}

	return ctx.JSON(http.StatusOK, theReservation)
}

func (srv *server) GetStatus(ctx echo.Context) error {
	username := ctx.Request().Header.Get("X-User-Name")
	status, err := srv.loyalty.GetStatus(username)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusOK, status)
}

func (srv *server) MakeReservation(ctx echo.Context) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	var model struct {
		HotelUID  string `json:"hotelUid"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}
	if err := json.Unmarshal(body, &model); err != nil {
		return fmt.Errorf("failed to unmarshal request body: %w", err)
	}
	hotelUID := model.HotelUID
	hotels, err := srv.reservation.GetAllHotels()
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	hotelExists := false
	var index int
	for index = range hotels {
		if hotels[index].HotelUID == hotelUID {
			hotelExists = true
			break
		}
	}
	if !hotelExists {
		return ctx.JSON(http.StatusBadRequest, echo.Map{})
	}

	dateLayout := "2006-01-02"
	startDate, err := time.Parse(dateLayout, model.StartDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	endDate, err := time.Parse(dateLayout, model.EndDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	duration := int(endDate.Sub(startDate).Hours() / 24)
	if duration < 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{})
	}

	username := ctx.Request().Header.Get("X-User-Name")
	user, err := srv.loyalty.GetUser(username)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	discount := user.Discount

	price := duration * hotels[index].Price * (100 - discount) / 100
	thePayment := payment.Payment{
		PaymentUID: uuid.New().String(),
		Status:     "PAID",
		Price:      price,
	}
	err = srv.payment.CreatePayment(thePayment)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}

	theReservation := reservation.Reservation{
		ReservationUID: uuid.New().String(),
		Username:       username,
		StartDate:      startDate,
		EndDate:        endDate,
		Status:         "PAID",
		HotelID:        hotels[index].ID,
		PaymentUID:     thePayment.PaymentUID,
	}

	err = srv.reservation.MakeReservation(theReservation)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}

	err = srv.loyalty.IncrementCounter(username)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}

	return ctx.JSON(http.StatusOK, echo.Map{})
}

func (srv *server) CancelReservation(ctx echo.Context) error {
	reservationUID := ctx.Param("reservationUid")
	err := srv.reservation.CancelReservation(reservationUID)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{})
	}

	reservation, err := srv.reservation.GetReservation(reservationUID)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{})
	}
	err = srv.payment.CancelPayment(reservation.PaymentUID)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{})
	}

	username := ctx.Request().Header.Get("X-User-Name")
	err = srv.loyalty.DecrementCounter(username)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, echo.Map{})
	}
	return nil
}

func (srv *server) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}
