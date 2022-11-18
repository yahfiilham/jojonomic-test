package models

type Rekening struct {
	ReffID string `json:"reff_id" gorm:"type:varchar(15);primaryKey"`
	RekeningData
	CreatedAt int64 `json:"created_at" gorm:"type:integer;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:integer;autoUpdateTime"`
}

type RekeningData struct {
	NoRek string  `json:"no_rek" gorm:"type:varchar(20);unique"`
	Saldo float64 `json:"saldo" gorm:"type:decimal(12,3)"`
}
