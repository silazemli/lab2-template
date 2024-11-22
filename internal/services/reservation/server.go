package reservation

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type server struct {
	srv echo.Echo
	rdb reservationStorage
	hdb hotelStorage
}

func NewServer(rdb reservationStorage, hdb hotelStorage) server {
	srv := server{}
	srv.rdb = rdb
	srv.hdb = hdb
	srv.srv = *echo.New()
	api := srv.srv.Group("api/v1")
	api.GET("/hotels", srv.GetAllHotels)
	api.GET("/reservations", srv.GetAllReservations)
	api.GET("/reservations/:reservationUID", srv.GetReservation)
	api.POST("/reservations", srv.MakeReservation)
	api.PATCH("/reservations/:reservationUID", srv.CancelReservation)

	return srv
}

func (srv *server) Start() error {
	err := srv.srv.Start("8070")
	if err != nil {
		return err
	}
	return nil
}

func (srv *server) GetAllReservations(ctx echo.Context) error {
	username := ctx.Request().Header.Get("X-User-Name")
	reservations, err := srv.rdb.GetReservations(username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusOK, reservations)
}

func (srv *server) GetAllHotels(ctx echo.Context) error {
	hotels, err := srv.hdb.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusOK, hotels)
}

func (srv *server) GetReservation(ctx echo.Context) error {
	reservation, err := srv.rdb.GetReservation(ctx.Param("reservationUid"))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.JSON(http.StatusNotFound, echo.Map{})
	}
	if reservation.Username != ctx.Request().Header.Get("X-User-Name") {
		return ctx.JSON(http.StatusForbidden, echo.Map{})
	}
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err})
	}

	return ctx.JSON(http.StatusOK, reservation)
}

func (srv *server) MakeReservation(ctx echo.Context) error {
	reservation := Reservation{}
	err := ctx.Bind(&reservation)
	reservation.Username = ctx.Request().Header.Get("X-User-Name")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	err = srv.rdb.MakeReservation(reservation)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusCreated, echo.Map{})
}

func (srv *server) CancelReservation(ctx echo.Context) error {
	reservationUID := ctx.Param("uid")
	err := srv.rdb.CancelReservation(reservationUID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusCreated, echo.Map{})
}
