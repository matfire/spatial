package types

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ConversionBody struct {
	From        string       `json:"from"`
	To          string       `json:"to"`
	Coordinates []Coordinate `json:"coordinates"`
}

type ConversionResult struct {
	Coordinates []Coordinate `json:"coordinates"`
}
