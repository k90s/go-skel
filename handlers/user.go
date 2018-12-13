package handlers

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/kai-ding/go-skel/domain"
	"github.com/kai-ding/go-skel/domain/impl"
	"github.com/labstack/echo"
)

type UserHandler struct {
	Service *impl.UserService
}

func (h *UserHandler) User(c echo.Context) error {
	id := 1
	user, err := h.Service.User(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := domain.User{
		Name: "test",
	}
	err := h.Service.CreateUser(&user)
	if err != nil {
		spew.Dump(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}
