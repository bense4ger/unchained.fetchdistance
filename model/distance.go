package model

import (
	"fmt"
	"time"
)

// Distance encapsulates the Length and Time from the starting location
type Distance struct {
	ID       string        `json:"id"`
	Length   string        `json:"length"`
	Duration time.Duration `json:"duration"`
}

func (d *Distance) String() string {
	return fmt.Sprintf("ID: %s Length: %s Duration: %d", d.ID, d.Length, d.Duration)
}
