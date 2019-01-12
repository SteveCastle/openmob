package cms

import "testing"

func TestSum(t *testing.T) {
	landingPage := GetAllLandingPages(5).Title
	if landingPage != "Hello World." {
		t.Errorf("Title was incorrect, got: %s, want: %s.", landingPage, "Hello World.")
	}
}
