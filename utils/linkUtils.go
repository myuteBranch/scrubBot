package utils

import (
	"fmt"
	"strings"
)

// Links map with useful links
var Links = map[string]string{
	"mhwi_deco_rates": "https://mhworld.kiranico.com/decorations",
	"mh_rage_reddit":  "https://www.reddit.com/r/monsterhunterrage/",
}

// GetLinkForMessage retunr link for message
func GetLinkForMessage(s string) string {
	if len(strings.Split(s, " ")) > 1 && strings.Split(s, " ")[1] != "" {
		if strings.Split(s, " ")[1] == "help" {
			keys := make([]string, 0, len(Links))
			for key := range Links {
				keys = append(keys, key)
			}

			return fmt.Sprintf("```Available links are: \n!linkMe %s ```", strings.Join(keys, "\n!linkMe "))
		}
		if Links[strings.Split(s, " ")[1]] != "" {

			return fmt.Sprintf("%s :  %s ", strings.Split(s, " ")[1], Links[strings.Split(s, " ")[1]])
		}
	}

	return "```Invalid argument for command !linkMe for valid options try \n try !linkMe help  ```"
}
