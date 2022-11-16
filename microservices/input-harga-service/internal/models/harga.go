package models

type Harga struct {
	ReffID string `json:"id"`
	HargaData
}

type HargaData struct {
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
	AdminID      string  `json:"admin_id"`
}
