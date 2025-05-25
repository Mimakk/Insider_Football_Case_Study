package simulator

type TeamRow struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
	GD     int    `json:"gd"`
	GF     int    `json:"gf"`
	GA     int    `json:"ga"`
}
