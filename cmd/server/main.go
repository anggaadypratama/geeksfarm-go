package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"geeksfarm-go/internal/config"
	"geeksfarm-go/internal/di"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	cfg, err := config.LoadConfig()

	if err != nil {
		e.Logger.Fatal(err)
	}

	db, err := config.SetupDB(cfg.Env.DatabaseURL)
	if err != nil {
		e.Logger.Fatal(err)
	}

	di.InitializeFeature(e, db)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, worled !")
	})
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(int(cfg.Env.Port))))
}
