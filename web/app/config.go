package app

import (
	"encoding/json"
	"time"
)

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
	Score  uint32           `json:"score"` // Realmente precisa?
	Colors map[string]Color `json:"colors"`
}

type BroadcastConfig struct {
	Address  string
	Insecure bool `json:"-"`
}

type Configuration struct {
	Broadcast         BroadcastConfig
	DevMode           bool              `json:"dev_mode"`
	StartMode         string            `json:"start_mode"`
	TimeRemaining     string            `json:"time_remaining"` // Realmente precisa?
	GameDuration      uint32            `json:"-"`
	ListeningDuration time.Duration     `json:"-"`
	HomeTeam          TeamConfiguration `json:"home_team"`
	AwayTeam          TeamConfiguration `json:"away_team"`
}