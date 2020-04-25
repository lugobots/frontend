package app

import "encoding/json"

type Color struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

func (c Color) MarshalJSON() ([]byte, error) {
	return json.Marshal([]int{
		c.R, c.G, c.B,
	})
}

type TeamConfiguration struct {
	Name   string           `json:"name"`
	Avatar string           `json:"avatar"`
	Colors map[string]Color `json:"colors"`
}

type Configuration struct {
	DevMode   bool              `json:"dev_mode"`
	StartMode string            `json:"start_mode"`
	HomeTeam  TeamConfiguration `json:"homeTeam"`
	AwayTeam  TeamConfiguration `json:"awayTeam"`
}
