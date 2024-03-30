package domain

type Element struct {
	Water rune `json:"water"`
	Wind  rune `json:"wind"`
}

type Status struct {
	WaterStatus string
	WindStatus  string
}
