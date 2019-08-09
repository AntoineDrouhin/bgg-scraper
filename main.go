package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	var topBoardGame [100]string
	var i = 0

	fmt.Println("Beginning")

	c := colly.NewCollector(
		colly.AllowedDomains("boardgamegeek.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	// c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	// 	fmt.Println("First column of a table row:", e.Text)
	// })

	c.OnHTML("tr td:nth-child(3)", func(e *colly.HTMLElement) {
		topBoardGame[i] = strings.TrimSpace(e.Text)
		i++
	})

	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://boardgamegeek.com/browse/boardgame")

	for i = 0; i < 100; i++ {
		fmt.Println("topBoardGame[", i, "]:", topBoardGame[i])
	}

}
