package app

import (
	"encoding/json"
	"net/http"

	"check-harga-service/internal/config"
	"check-harga-service/internal/models"
)

func CheckHarga(w http.ResponseWriter, r *http.Request) {
	c := config.NewConfig()

	m := new(models.Harga)

	if err := c.DB.Model(m).First(m).Error; err != nil {
		rs := &models.Response{
			Error:   true,
			Message: err.Error(),
		}
		sendResponse(w, http.StatusBadRequest, rs)
		return
	}

	rs := &models.Response{
		Error: false,
		Data:  m,
	}
	sendResponse(w, http.StatusOK, rs)
}

func sendResponse(w http.ResponseWriter, code int, rs *models.Response) {
	resp, err := json.Marshal(rs)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
