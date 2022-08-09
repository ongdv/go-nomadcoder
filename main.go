package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ongdv/learngo/scrapper"
)

func handleHome(c echo.Context) error {
	return c.File("./html/index.html")
}

func handleScrape(c echo.Context) error {
	term := scrapper.CleanStrings(c.FormValue("terms"))
	scrapper.Scrape(term)
	return c.File("/html/index.html")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":8080"))
}
