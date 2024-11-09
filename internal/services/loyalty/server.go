package loyalty

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type server struct {
	srv echo.Echo
	db loyaltyStorage
}

func NewServer(db loyaltyStorage) server {
	srv := server{}
	srv.db = db
	srv.srv = *echo.New()
	api := srv.srv.Group("/api/loyalty")
	api.GET("/username", srv.GetStatus)
	api.PATCH("/increment/username", srv.IncrementCounter)
	api.PATCH("/decrement/username", srv.DecrementCounter)
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
	if errors.Is(err, )
}

func (srv *server) IncrementCounter(ctx echo.Context) error {

}

func (srv *server) DecrementCounter(ctx echo.Context) error {

}