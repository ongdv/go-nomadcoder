package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/ongdv/learngo/scrapper"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("./html/index.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := scrapper.CleanStrings(c.FormValue("term"))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":8080"))
}
