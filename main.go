package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// NASAImage represents the data structure for scraped data
type NASAImage struct {
	Src  string `json:"src"`
	Alt  string `json:"alt"`
	Date string `json:"date"`
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	var nasaImage NASAImage

	c.OnHTML("img", func(e *colly.HTMLElement) {
		nasaImage.Src = "https://apod.nasa.gov/apod/" + e.Attr("src")
		nasaImage.Alt = e.Attr("alt")
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		nasaImage.Date = e.ChildText("p:nth-child(3)")
	})

	c.Visit("https://apod.nasa.gov/apod/astropix.html")

	// Serialise the data to JSON
	jsonData, err := json.Marshal(nasaImage)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save the data to a file
	outputFile, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating output file", err)
		return
	}

	defer outputFile.Close()

	_, err = outputFile.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to output file", err)
		return
	}

	// Format a nice human readable output of all the data scrapped and the json data serialized
	fmt.Println("\n\x1b[31mNASA Image Information:\x1b[0m")
	fmt.Printf("\x1b[33mSrc:\x1b[0m \x1b[34m%s\x1b[0m\n", nasaImage.Src)
	fmt.Printf("\x1b[33mAlt:\x1b[0m %s\n", nasaImage.Alt)
	fmt.Printf("\x1b[33mDate:\x1b[0m %s\n\n", nasaImage.Date)
	fmt.Printf("\x1b[32mJSON Data:\x1b[0m %s\n", string(jsonData))
}
