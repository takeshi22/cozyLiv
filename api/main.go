package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/takeshi22/cozyLiv/infrastructure"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Are you OK?")
	})

	// e.GET("/read_csv", func(c echo.Context) error {
	// 	entities := helper.ReadCsv()

	// 	infrastructure.Db.Create(*&entities)

	// 	return nil
	// })

	err := infrastructure.InitDatabase()
	if err != nil {
		e.Logger.Fatal("could not create database", err)
	}
	e.Logger.Fatal(e.Start(":1313"))

}
