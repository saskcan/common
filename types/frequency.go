package types

type Frequency int

const (
	MINUTE Frequency = iota
	HOUR
	DAY
	WEEK
	MONTH
	YEAR
)
