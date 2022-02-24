package main

import (
	"github.com/labstack/echo"
	"job-scrapper/scrapper"
	"os"
	"strings"
)

const FileName string = "jobs.csv"

// Handler
func hello(c echo.Context) error {
	return c.File("home.html")
}

// Handler
func handleScrape(c echo.Context) error {
	defer os.Remove(FileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(FileName, FileName)
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
