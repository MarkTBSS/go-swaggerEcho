package main

import (
	"net/http"

	_ "github.com/MarkTBSS/go-swaggerEcho/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/pets", allPet)
	e.Logger.Fatal(e.Start(":1323"))
}

type Pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type HTTPError struct {
	BusinessErrorCode int    `json:"businessErrorCode"`
	Description       string `json:"description"`
}

// ListPet godoc
// @Tags         pet
// @Summary      List of all pet
// @Router       /pets [get]
func allPet(c echo.Context) error {
	return c.JSON(http.StatusOK, []Pet{
		{ID: 1, Name: "ChaNom"},
		{ID: 2, Name: "Mali"},
	})
}
