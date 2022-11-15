package models

type Harga struct {
	ID string `json:"id" gorm:"primaryKey"`
	HargaData
	CreatedAt int64 `json:"created_at" gorm:"type:int;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:int;autoUpdateTime"`
}

type HargaData struct {
	HargaTopup   int64 `json:"harga_topup"`
	HargaBuyback int64 `json:"harga_buyback"`
}
