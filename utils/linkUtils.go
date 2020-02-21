package utils

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// GetLinkForMessage retunr link for message <word> reply with link for <word> if exist
func GetLinkForMessage(s string) string {
	if len(strings.Split(s, " ")) > 1 && strings.Split(s, " ")[1] != "" {
		if strings.Split(s, " ")[1] == "help" {
			links := ReadAllLinks()
			keys := make([]string, 0, len(links))
			for _, link := range links {
				keys = append(keys, link.Descrip)
			}

			return fmt.Sprintf("```Available links are: \n!link %s ```", strings.Join(keys, "\n!link "))
		}
		links := ReadAllLinks(strings.Split(s, " ")[1])
		if len(links) > 0 {
			retString := "Heres what I found!\n"
			for _, link := range links {
				retString += fmt.Sprintf("%s :  %s \n", link.Descrip, link.Link)
			}
			return retString
		}
	}
	return "```Invalid argument for command !link for valid options try \n try !link help  ```"
}

// ReadAllLinks used for testing
func ReadAllLinks(filter ...string) []*LinkObj {
	Log.Warnln(filter)
	db := GetDbConnection("sqlite3", "bot.db")
	defer db.Close()
	// rows, er := db.Queryx("select * from link_store")
	links := []*LinkObj{}
	var er error
	if len(filter) == 0 {
		er = db.Select(&links, "select * from link_store")
	} else {
		er = db.Select(&links, "select * from link_store where descrip like $1", "%"+filter[0]+"%")
	}

	if er != nil {
		Log.Errorln(errors.Wrap(er, "Unable to fetch link_store"))
	}
	Log.Warnln(links)
	for _, link := range links {
		Log.Infoln(link.Username)
		Log.Infoln(link.Descrip)
		Log.Infoln(link.Link)
		Log.Infoln(link.ID)
	}
	return links
}

//Addlink is to add link from bot
func Addlink(user string, s string) string {
	if len(strings.Split(s, " ")) > 1 && strings.Split(s, " ")[2] != "" && strings.Split(s, " ")[3] != "" {
		if addLinkToDb(user, strings.Split(s, " ")[2], strings.Split(s, " ")[3]) {
			return fmt.Sprintf("Successfully stored link %v with alias %v", strings.Split(s, " ")[3], strings.Split(s, " ")[2])
		}
	}
	return "Failed to store Link please use this format `!link add <alias> <link> ` "
}

func addLinkToDb(username string, descrip string, link string) bool {
	db := GetDbConnection("sqlite3", "bot.db")
	defer db.Close()
	insertQuery := `INSERT INTO link_store (username, link, descrip) VALUES(?, ?, ?)`
	_, err := (db.MustExec(insertQuery, username, link, descrip)).RowsAffected()
	if err != nil {
		Log.Errorln(err)
		return false
	}
	return true
}
