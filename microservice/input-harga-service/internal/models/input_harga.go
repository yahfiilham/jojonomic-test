package models

type Harga struct {
	ReffID string `json:"id"`
	HargaData
}

type HargaData struct {
	HargaTopup   int64  `json:"harga_topup"`
	HargaBuyback int64  `json:"harga_buyback"`
	AdminID      string `json:"admin_id"`
}
