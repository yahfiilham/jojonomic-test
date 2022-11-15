package models

type Harga struct {
	ID string `json:"id" gorm:"type:varchar(255);primaryKey"`
	HargaData
	CreatedAt int64 `json:"created_at" gorm:"type:integer;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:integer;autoUpdateTime"`
}

type HargaData struct {
	HargaTopup   int64  `json:"harga_topup" gorm:"type:integer;"`
	HargaBuyback int64  `json:"harga_buyback" gorm:"type:integer"`
	AdminID      string `json:"admin_id" gorm:"type:varchar(255);"`
}
