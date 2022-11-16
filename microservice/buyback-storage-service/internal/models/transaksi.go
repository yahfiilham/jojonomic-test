package models

type Transaksi struct {
	ReffID string `json:"reff_id" gorm:"type:varchar(15);primaryKey"`
	TransaksiData
	CreatedAt int64 `json:"created_at" gorm:"type:integer;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:integer;autoUpdateTime"`
}

type TransaksiData struct {
	Type         string  `json:"type" gorm:"type:varchar(20)"`
	Gram         float64 `json:"gram" gorm:"type:decimal(12,2)"`
	Saldo        float64 `json:"saldo" gorm:"type:decimal(12,2)"`
	HargaTopup   float64 `json:"harga_topup" gorm:"type:integer;"`
	HargaBuyback float64 `json:"harga_buyback" gorm:"type:integer"`
	NoRek        string  `json:"no_rek" gorm:"type:varchar(15)"`
}
