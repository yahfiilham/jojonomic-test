package models

type Harga struct {
	ID string `json:"id"`
	HargaData
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type HargaData struct {
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
	AdminID      string  `json:"admin_id"`
}
