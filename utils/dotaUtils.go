package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Match type
type Match struct {
	LeftTeam  string `json:"LeftTeam"`
	RightTeam string `json:"RightTeam"`
	TimeStamp string `json:"TimeStamp"`
}

func getDataFromFile() (map[string][]Match, error) {
	tourneyMap := make(map[string][]Match)
	f, fErr := os.Open("./matches.json")
	if fErr != nil {
		return tourneyMap, errors.New("file does not exist")
	}
	fi, _ := f.Stat()
	t1 := time.Now()
	Log.Debug(fi.ModTime())
	Log.Debug(t1.Sub(fi.ModTime()).Hours())
	if t1.Sub(fi.ModTime()).Hours() > 24.0 {
		return tourneyMap, errors.New("file is too old")
	}
	byteValue, _ := ioutil.ReadAll(f)
	em := json.Unmarshal(byteValue, &tourneyMap)

	if em != nil {
		panic(em)
	}
	return tourneyMap, nil
}

func writeToFile(data map[string][]Match) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// sanity check
	// Log.Info(string(jsonData))

	// write to JSON file

	jsonFile, err := os.Create("./matches.json")

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	Log.Info("JSON data written to ", jsonFile.Name())

}

// GetDotaMatches returns a set of web scapred matches
func GetDotaMatches() map[string][]Match {
	data, err := getDataFromFile()
	if err == nil {
		Log.Info("Loading data from local file")
		return data
	}
	Log.Info(err)

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
	writeToFile(tourneyMap)

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
