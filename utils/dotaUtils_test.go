package utils

import (
	"strings"
	"testing"
)

func TestGetFormatedMatches(t *testing.T) {
	res := GetFormatedMatches(map[string][]Match{
		"test": {
			Match{
				LeftTeam:  "eg",
				RightTeam: "og",
				TimeStamp: "002",
			},
			Match{
				LeftTeam:  "reg",
				RightTeam: "geg",
				TimeStamp: "001",
			},
		},
	})

	if strings.TrimSpace(strings.Split(res, "\n")[1]) != "test" || strings.TrimSpace(strings.Split(res, "\n")[2]) != "eg vs og  | at 002" {
		t.Errorf("format incorrect, got: %s, want: %s.", strings.Split(res, "\n")[1:3], []string{"test", " eg vs og  | at 002 "})
	}
}
