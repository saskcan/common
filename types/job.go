package types

type Job struct {
	Symbol    Symbol    `json:"symbol"`
	Range     DateRange `json:"date_range"`
	Frequency Frequency `json:"frequency"`
}
