package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"check-saldo-service/configs"
	"check-saldo-service/internal/models"

	"gorm.io/gorm"
)

func CheckSaldo(w http.ResponseWriter, r *http.Request) {
	c := configs.NewConfig()

	p := new(models.CheckSaldoRequest)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(p); err != nil {
		sendResponse(w, http.StatusBadRequest, true, ``, err.Error(), nil)
		return
	}

	m := new(models.Rekening)

	if err := c.DB.Model(m).Where("no_rek = ?", p.Norek).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendResponse(w, http.StatusBadRequest, true, ``, `rekening tidak ditemukan`, nil)
			return
		}
		sendResponse(w, http.StatusBadRequest, true, ``, err.Error(), nil)
		return
	}

	sendResponse(w, http.StatusOK, false, m.ReffID, ``, m)
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
