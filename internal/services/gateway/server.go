package gateway

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/silazemli/lab2-template/internal/services/gateway/clients"
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
	client := &http.Client{
		Transport: &http.Transport{MaxConnsPerHost: 10},
		Timeout:   100,
	}
	srv.loyalty = *clients.NewLoyaltyClient(client, srv.cfg.LoyaltyService)
	srv.payment = *clients.NewPaymentClient(client, srv.cfg.PaymentService)
	srv.reservation = *clients.NewReservationClient(client, srv.cfg.ReservationService)

	api := srv.srv.Group("/api/v1")
	api.GET("/hotels", srv.GetAllHotels)
	api.GET("/me", srv.GetUser)
	api.GET("/loyalty", srv.GetStatus)
	api.GET("/reservations", srv.GetAllReservations)
	api.GET("reservations/:reservationUid", srv.GetReservation)
	api.POST("/reservations", srv.MakeReservation)
	api.DELETE("/reservations/:reservationUid", srv.CancelReservation)
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
	
}

func (srv *server) GetAllHotels(ctx echo.Context) error {

}

func (srv *server) GetAllReservations(ctx echo.Context) error {

}

func (srv *server) GetReservation(ctx echo.Context) error {

}

func (srv *server) GetStatus(ctx echo.Context) error {

}

func (srv *server) MakeReservation(ctx echo.Context) error {

}

func (srv *server) CancelReservation(ctx echo.Context) error {

}
