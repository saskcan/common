package types

import "time"

// Job represents a request for work to be done
type Job struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// CandlesSinceDateRequest represents a request to retrieve candles since a particular date for a given symbol
type CandlesSinceDateRequest struct {
	Symbol    Symbol    `json:"symbol"`
	Exchange  string    `json:"exchange"`
	Frequency string    `json:"frequency"`
	Date      time.Time `json:"date"`
}
