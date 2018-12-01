package main

import (
	"database/sql"
	"gobase/http"
	"gobase/impl"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	db, err := sql.Open("postgresql", "postgres://")
	if err != nil {
		panic(err)
	}

	userService := &impl.UserService{db}
	userHandler := &http.UserHandler{userService}

	user := e.Group("/users")
	user.GET("/:id", userHandler.User)
}
