package handler

import (
	"lemonilo/store"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	userStore store.UserStore
}

func NewHandler(db *gorm.DB, c *echo.Echo) *Handler {
	return &Handler{
		userStore: store.NewUserStore(db),
	}
}
