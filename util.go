package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type match struct {
	LeftTeam  string `json:"LeftTeam"`
	RightTeam string `json:"RightTeam"`
	TimeStamp string `json:"TimeStamp"`
}

var links = map[string]string{
	"mhwi_deco_rates": "https://mhworld.kiranico.com/decorations",
	"mh_rage_reddit":  "https://www.reddit.com/r/monsterhunterrage/",
}

func getDotaMatches() map[string][]match {
	tourneyMap := make(map[string][]match)

	// Make HTTP request
	response, err := http.Get("https://liquipedia.net/dota2/Liquipedia:Upcoming_and_ongoing_matches")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".infobox_matches_content").Each(func(i int, s *goquery.Selection) {

		leftTeam := strings.TrimSpace(s.Find(".team-left").Text())
		rightTeam := strings.TrimSpace(s.Find(".team-right").Text())
		timeStamp := strings.TrimSpace(s.Find(".timer-object-countdown-only").Text())
		tourney := strings.TrimSpace(s.Find("a").Last().Text())

		tourneyMap[tourney] = append(tourneyMap[tourney], match{
			LeftTeam:  leftTeam,
			RightTeam: rightTeam,
			TimeStamp: timeStamp,
		})

	})
	return tourneyMap
}

func getFormatedMatches(tourneyMap map[string][]match) string {
	returnString := "Upcoming Dota Matches \n"
	for tourney, matches := range tourneyMap {
		returnString += fmt.Sprintf("%s \n", tourney)
		for _, match := range matches {
			returnString += fmt.Sprintf(" %s vs %s  | at %s \n", match.LeftTeam, match.RightTeam, match.TimeStamp)
		}

	}
	return returnString
}

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}
