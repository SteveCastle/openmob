package geography

// Point is a location on a globe.
type Point struct {
	Lat float64
	Lng float64
}

// Region is an enclosed shape on a globe.
type Region []Point

// NullRegion is an optional enclosed shape on a globe.
type NullRegion struct {
	Valid  bool
	Region Region
}
