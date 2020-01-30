package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Match type
type Match struct {
	LeftTeam  string `json:"LeftTeam"`
	RightTeam string `json:"RightTeam"`
	TimeStamp string `json:"TimeStamp"`
}

// GetDotaMatches returns a set of web scapred matches
func GetDotaMatches() map[string][]Match {
	tourneyMap := make(map[string][]Match)

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

		tourneyMap[tourney] = append(tourneyMap[tourney], Match{
			LeftTeam:  leftTeam,
			RightTeam: rightTeam,
			TimeStamp: timeStamp,
		})

	})
	return tourneyMap
}

// GetFormatedMatches returns a formated string of matches
func GetFormatedMatches(tourneyMap map[string][]Match) string {
	returnString := "Upcoming Dota Matches \n"
	for tourney, matches := range tourneyMap {
		returnString += fmt.Sprintf("%s \n", tourney)
		for _, match := range matches {
			returnString += fmt.Sprintf(" %s vs %s  | at %s \n", match.LeftTeam, match.RightTeam, match.TimeStamp)
		}

	}
	return returnString
}

// ChunkString chunks a string to a given size
func ChunkString(s string, chunkSize int) []string {
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
