package main

import (
	"fmt"
	"log"

	"bggscraper/scraper"

	"github.com/gocolly/colly"
)

func main() {

	var topBoardGame map[int]scraper.BoardGame = make(map[int]scraper.BoardGame)

	var i = 1

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

	c.OnHTML("tr td:nth-child(3)", func(e *colly.HTMLElement) {
		newBoardGame := scraper.BoardGame{Rank: i}
		newBoardGame.Name, newBoardGame.ReleaseDate = scraper.ExtractGameName(e.Text)

		topBoardGame[i] = newBoardGame
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
