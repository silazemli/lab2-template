package payment

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type server struct {
	srv echo.Echo
	db  paymentStorage
}

func NewServer(db paymentStorage) server {
	srv := server{}
	srv.db = db
	srv.srv = *echo.New()
	api := srv.srv.Group("/api/v1")
	api.POST("", srv.PostPayment)
	api.PATCH("/:uid", srv.CancelPayment)

	return srv
}

func (srv *server) Start() error {
	err := srv.srv.Start(":8060")
	if err != nil {
		return err
	}
	return nil
}

func (srv *server) PostPayment(ctx echo.Context) error {
	var thePayment Payment
	err := ctx.Bind(&thePayment)
	if err != nil {
		return err
	}
	err = srv.db.PostPayment(thePayment)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, echo.Map{})
}

func (srv *server) CancelPayment(ctx echo.Context) error {
	uid := ctx.Param("uid")
	err := srv.db.CancelPayment(uid)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{})
}
