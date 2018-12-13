package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/kai-ding/go-skel/config"
	"github.com/kai-ding/go-skel/domain/impl"
	"github.com/kai-ding/go-skel/handlers"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Debug = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "fuck!!!")
	})

	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.C.Postgres.User,
		config.C.Postgres.Password,
		config.C.Postgres.Host,
		config.C.Postgres.Port,
		config.C.Postgres.DBName,
		config.C.Postgres.SSLMode,
	))
	if err != nil {
		panic(err)
	}

	userService := &impl.UserService{db}
	userHandler := &handlers.UserHandler{userService}

	user := e.Group("/users")
	user.GET("/:id", userHandler.User)
	user.POST("/", userHandler.CreateUser)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.C.Server.Port)))
}
