package models

type Harga struct {
	ID string `json:"id"`
	HargaData
}

type HargaData struct {
	AdminID      string `json:"admin_id"`
	HargaTopup   int    `json:"harga_topup"`
	HargaBuyback int    `json:"harga_buyback"`
}
