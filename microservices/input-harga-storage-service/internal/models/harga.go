package models

type Harga struct {
	ReffID string `json:"reff_id" gorm:"type:varchar(15);primaryKey"`
	HargaData
	CreatedAt int64 `json:"created_at" gorm:"type:integer;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:integer;autoUpdateTime"`
}

type HargaData struct {
	HargaTopup   float64 `json:"harga_topup" gorm:"type:decimal(12,3);"`
	HargaBuyback float64 `json:"harga_buyback" gorm:"type:decimal(12,3)"`
	AdminID      string  `json:"admin_id" gorm:"type:varchar(15);"`
}
