package utils

import (
	"testing"
)

func TestGetLinkForMessage(t *testing.T) {
	link := GetLinkForMessage("!linkMe mhwi_deco_rates")
	if link != "mhwi_deco_rates :  https://mhworld.kiranico.com/decorations " {
		t.Errorf("link incorrect, got: %s, want: %s.", link, "mhwi_deco_rates :  https://mhworld.kiranico.com/decorations")
	}
}
