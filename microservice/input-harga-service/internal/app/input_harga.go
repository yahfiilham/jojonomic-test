package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"input-harga-service/internal/config"
	"input-harga-service/internal/models"

	"github.com/segmentio/kafka-go"
	"github.com/teris-io/shortid"
)

func InputHarga(w http.ResponseWriter, r *http.Request) {
	c := config.NewConfig()
	var data models.HargaData

	id, _ := shortid.Generate()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		rs := &models.Response{
			Error:   true,
			ReffID:  id,
			Message: err.Error(),
		}
		sendResponse(w, http.StatusBadRequest, rs)
		return
	}

	m := &models.Harga{
		ID:        id,
		HargaData: data,
	}

	defer r.Body.Close()

	pByte, err := json.Marshal(&m)
	if err != nil {
		rs := &models.Response{
			Error:   true,
			ReffID:  id,
			Message: err.Error(),
		}
		sendResponse(w, http.StatusBadRequest, rs)
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
		rs := &models.Response{
			Error:   true,
			ReffID:  id,
			Message: "Kafka not ready",
		}
		sendResponse(w, http.StatusBadRequest, rs)
		return
	}

	// send response
	rs := &models.Response{
		Error:  false,
		ReffID: id,
	}
	sendResponse(w, http.StatusCreated, rs)
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
