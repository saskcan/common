package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Candle struct {
	ProductID uuid.UUID `json:product_id`
	Open      float32   `json:"open"`
	High      float32   `json:"high"`
	Low       float32   `json:"low"`
	Close     float32   `json:"close"`
	StartTime time.Time `json:"start_time"`
	Frequency Frequency `json:"frequency"`
}
