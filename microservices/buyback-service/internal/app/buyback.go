package app

import (
	"buyback-service/internal/config"
	"buyback-service/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

func Buyback(w http.ResponseWriter, r *http.Request) {
	c := config.NewConfig()
	p := new(models.BuybackRequest)

	id, _ := shortid.Generate()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	rek, err := getRekening(c.DB, p.Norek)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	harga, err := getCurrentHarga(c.DB)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	if p.Gram > rek.Saldo {
		sendResponse(w, http.StatusBadRequest, true, id, "saldo emas anda tidak mencukupi")
		return
	}

	m := &models.BuybackData{
		ReffID:       id,
		Gram:         p.Gram,
		Harga:        p.Harga,
		Norek:        p.Norek,
		Saldo:        rek.Saldo,
		HargaTopup:   harga.HargaTopup,
		HargaBuyback: harga.HargaBuyback,
	}

	defer r.Body.Close()

	pByte, err := json.Marshal(&m)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	c.Kafka.SetWriteDeadline(time.Now().Add(10 * time.Second))

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("address-%s", r.RemoteAddr)),
		Value: pByte,
	}

	_, err = c.Kafka.WriteMessages(msg)
	if err != nil {
		log.Printf("error write message kafka with error message : %s\n", err.Error())
		sendResponse(w, http.StatusBadRequest, true, id, "Kafka not ready")
		return
	}

	sendResponse(w, http.StatusCreated, false, id, ``)
}

func getRekening(db *gorm.DB, noRek string) (*models.Rekening, error) {
	var rek *models.Rekening
	if err := db.Model(rek).Where("no_rek = ?", noRek).First(&rek).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rekening tidak ditemukan")
		}
		return nil, err
	}

	return rek, nil
}

func getCurrentHarga(db *gorm.DB) (*models.Harga, error) {
	var h *models.Harga
	if err := db.Model(h).Order("created_at DESC").First(&h).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("harga tidak ditemukan")
		}
		return nil, err
	}

	return h, nil
}

func sendResponse(w http.ResponseWriter, code int, isError bool, id string, msg string) {
	rs := &models.Response{
		Error:   isError,
		ReffID:  id,
		Message: msg,
	}

	resp, err := json.Marshal(rs)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
