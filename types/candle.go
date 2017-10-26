package types

import (
	"time"
)

type Candle struct {
	Symbol    Symbol    `json:symbol`
	Open      float32   `json:"open"`
	High      float32   `json:"high"`
	Low       float32   `json:"low"`
	Close     float32   `json:"close"`
	StartTime time.Time `json:"start_time"`
	Frequency Frequency `json:"frequency"`
}
