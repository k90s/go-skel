package http

import (
	"net/http"

	skel "github.com/kai-ding/go-skel"
	"github.com/labstack/echo"
)

type UserHandler struct {
	S skel.UserService
}

func (h *UserHandler) User(c echo.Context) error {
	id := 1
	user, err := h.S.User(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
