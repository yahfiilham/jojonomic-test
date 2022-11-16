package models

type BuybackData struct {
	ReffID       string  `json:"reff_id"`
	Gram         float64 `json:"gram"`
	Harga        float64 `json:"harga"`
	Norek        string  `json:"norek"`
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
	Saldo        float64 `json:"saldo"`
}
