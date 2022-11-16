package models

type BuybackRequest struct {
	Gram  float64 `json:"gram"`
	Harga float64 `json:"harga"`
	Norek string  `json:"norek"`
}
