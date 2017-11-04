package types

// Product represents a product
type Product struct {
	ID       string `json:"id"`
	Symbol   Symbol `json:"symbol"`
	Exchange string `json:"exchange"`
}
