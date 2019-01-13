package cause

// Cause is the primary type for a cause.
type Cause struct {
	Title string
}

// GetAllCauses Returns all the causes.
func GetAllCauses(c int) Cause {
	return Cause{Title: "Hello World."}
}
