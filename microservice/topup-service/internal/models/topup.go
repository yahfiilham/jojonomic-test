package models

type Topup struct {
	ReffID string `json:"reff_id"`

	TopupData
}

type TopupData struct {
	Gram  string `json:"gram"`
	Harga string `json:"harga"`
	Norek string `json:"norek"`
}
