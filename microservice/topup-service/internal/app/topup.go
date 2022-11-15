package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"topup-service/internal/config"
	"topup-service/internal/models"

	"github.com/segmentio/kafka-go"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

func Topup(w http.ResponseWriter, r *http.Request) {
	c := config.NewConfig()
	var data models.TopupData

	id, _ := shortid.Generate()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	m := &models.Topup{
		ReffID:    id,
		TopupData: data,
	}

	defer r.Body.Close()

	pByte, err := json.Marshal(&m)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, true, id, err.Error())
		return
	}

	if err := validateTopup(c.DB, &m.TopupData); err != nil {
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

func validateTopup(db *gorm.DB, data *models.TopupData) error {
	m := new(models.Harga)

	if err := db.Model(m).First(&m).Error; err != nil {
		return err
	}

	harga, err := strconv.Atoi(data.Harga)
	if err != nil {
		return err
	}

	if int64(harga) != m.HargaData.HargaTopup {
		return errors.New("harga doesn't match with current harga topup")
	}

	gram, err := strconv.ParseFloat(data.Gram, 64)
	if err != nil {
		return err
	}

	n := gram / 0.001
	if n != float64(int(n)) || gram == 0 {
		return errors.New("minimum top-up is multiply of 0.001")
	}

	return nil
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
