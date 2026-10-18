package model

import "time"

type Status string

const (
	StatusScheduled Status = "scheduled"
	StatusLive      Status = "live"
	StatusFinal     Status = "final"
)

type Team struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type Game struct {
	ID        string    `json:"id"`
	League    string    `json:"league"`
	Home      Team      `json:"home"`
	Away      Team      `json:"away"`
	HomeScore int       `json:"homeScore"`
	AwayScore int       `json:"awayScore"`
	Status    Status    `json:"status"`
	Clock     string    `json:"clock"`
	StartTime time.Time `json:"startTime"`
}
