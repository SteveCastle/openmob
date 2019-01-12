// Package cms implements content management features.
package cms

// LandingPage is a visitor facing with a goal or goals.
type LandingPage struct {
	Title string
}

// GetAllLandingPages Returns all the landing pages for a cause.
func GetAllLandingPages(c int) LandingPage {
	return LandingPage{Title: "Hello World."}
}
