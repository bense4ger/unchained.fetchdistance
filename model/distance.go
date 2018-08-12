package model

import (
	"time"
)

// Distance encapsulates the Length and Time from the starting location
type Distance struct {
	ID       string        `json:"id"`
	Length   string        `json:"length"`
	Duration time.Duration `json:"duration"`
}
