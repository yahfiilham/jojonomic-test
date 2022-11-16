package models

type CheckMutasiRequest struct {
	Norek     string `json:"norek"`
	StartDate int32  `json:"start_date"`
	EndDate   int32  `json:"end_date"`
}
