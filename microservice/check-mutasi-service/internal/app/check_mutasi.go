package app

import (
	"check-mutasi-service/internal/config"
	"check-mutasi-service/internal/models"
	"encoding/json"
	"net/http"
)

func CheckMutasi(w http.ResponseWriter, r *http.Request) {
	c := config.NewConfig()

	p := new(models.CheckMutasiRequest)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(p); err != nil {
		sendResponse(w, http.StatusBadRequest, true, ``, err.Error(), nil)
		return
	}

	m := make([]*models.Transaksi, 0)

	if err := c.DB.Model(m).Where("no_rek = ? AND created_at >= ? AND created_at <= ?", p.Norek, p.StartDate, p.EndDate).Find(&m).Error; err != nil {
		sendResponse(w, http.StatusBadRequest, true, ``, err.Error(), nil)
		return
	}

	if len(m) == 0 {
		sendResponse(w, http.StatusNotFound, true, ``, "tidak ada data transaksi anda", nil)
		return
	}

	sendResponse(w, http.StatusOK, false, ``, ``, m)
}

func sendResponse(w http.ResponseWriter, code int, isError bool, id string, msg string, data interface{}) {
	rs := &models.Response{
		Error:   isError,
		ReffID:  id,
		Message: msg,
		Data:    data,
	}

	resp, err := json.Marshal(rs)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
