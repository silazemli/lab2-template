package loyalty

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type server struct {
	srv echo.Echo
	db  loyaltyStorage
}

func NewServer(db loyaltyStorage) server {
	srv := server{}
	srv.db = db
	srv.srv = *echo.New()
	api := srv.srv.Group("/api/loyalty")
	api.GET("/username", srv.GetStatus)
	api.PATCH("/increment/username", srv.IncrementCounter)
	api.PATCH("/decrement/username", srv.DecrementCounter)

	return srv
}

func (srv *server) Start() error {
	err := srv.srv.Start(":8050")
	if err != nil {
		return err
	}
	return nil
}

func (srv *server) GetStatus(ctx echo.Context) error {
	status, err := srv.db.GetStatus(ctx.Param("username"))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.JSON(http.StatusNotFound, echo.Map{})
	}
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusOK, status)
}

func (srv *server) IncrementCounter(ctx echo.Context) error {
	err := srv.db.IncrementCounter(ctx.Param("username"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{})
	}
	return nil
}

func (srv *server) DecrementCounter(ctx echo.Context) error {
	err := srv.db.DecrementCounter(ctx.Param("username"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{})
	}
	return nil
}
