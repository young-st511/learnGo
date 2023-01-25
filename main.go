package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/young-st511/learnGo/scrapper"
	// "github.com/young-st511/learnGo/scrapper"
)

const fileName = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))

	// scrapper.Scrape("프론트엔드")
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	searchTerm := c.FormValue("search")
	searchTerm = scrapper.CleanString(searchTerm)

	scrapper.Scrape(searchTerm)

	defer os.Remove(fileName)

	return c.Attachment(fileName, "job.csv")
}
