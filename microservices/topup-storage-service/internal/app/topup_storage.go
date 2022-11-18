package app

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"topup-storage-service/configs"
	"topup-storage-service/internal/models"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

func ReadMessage() {
	c := configs.NewConfig()
	ctx := context.Background()

	for {
		kf, err := c.Kafka.FetchMessage(ctx)
		if err != nil {
			break
		}
		log.Printf("message at topic/partition/offset %v/%v/%v: %s\n", kf.Topic, kf.Partition, kf.Offset, string(kf.Key))

		m := new(models.Topup)
		if err := json.Unmarshal(kf.Value, &m); err != nil {
			log.Printf("unmarshall data error : %s", err.Error())
		}

		if err := c.DB.Model(m).Create(m).Error; err != nil {
			log.Printf("error when insert topup to database with error message %s\n", err.Error())
		}

		saveToRekening(c.DB, m)
		saveToTransaksi(c.DB, m)

		if err := c.Kafka.CommitMessages(ctx, kf); err != nil {
			log.Printf("CommitMessage Failed error : %s", err.Error())
		}
	}
}

func saveToRekening(db *gorm.DB, data *models.Topup) {
	rek := getRekening(db, data.Norek)

	gram, err := strconv.ParseFloat(data.Gram, 64)
	if err != nil {
		log.Printf("error when parse harga with error message %s\n", err.Error())
	}

	if rek.ReffID != `` {
		saldo := rek.Saldo + gram
		rek.Saldo = saldo

		if err := db.Model(rek).Where("reff_id = ?", rek.ReffID).Updates(&rek).Error; err != nil {
			log.Printf("error when insert rekening to database with error message %s\n", err.Error())
		}
	}

	if rek.ReffID == `` {
		id, _ := shortid.Generate()
		r := &models.Rekening{
			ReffID: id,
			RekeningData: models.RekeningData{
				NoRek: data.Norek,
				Saldo: gram,
			},
		}

		if err := db.Model(r).Save(&r).Error; err != nil {
			log.Printf("error when insert rekening to database with error message %s\n", err.Error())
		}
	}
}

func saveToTransaksi(db *gorm.DB, data *models.Topup) {
	rek := getRekening(db, data.Norek)
	harga := getCurrentHarga(db)

	gram, err := strconv.ParseFloat(data.Gram, 64)
	if err != nil {
		log.Printf("error when parse harga with error message %s\n", err.Error())
	}

	id, _ := shortid.Generate()
	m := &models.Transaksi{
		ReffID: id,
		TransaksiData: models.TransaksiData{
			Type:         "topup",
			Gram:         gram,
			Saldo:        rek.Saldo,
			HargaTopup:   harga.HargaTopup,
			HargaBuyback: harga.HargaBuyback,
			NoRek:        data.Norek,
		},
	}

	if err := db.Model(m).Save(&m).Error; err != nil {
		log.Printf("error when insert transaction to database with error message %s\n", err.Error())
	}
}

func getRekening(db *gorm.DB, noRek string) *models.Rekening {
	var rek *models.Rekening
	if err := db.Model(rek).Where("no_rek = ?", noRek).First(&rek).Error; err != nil {
		log.Printf("error when get data rekening from database with error message %s\n", err.Error())
	}

	return rek
}

func getCurrentHarga(db *gorm.DB) *models.Harga {
	var h *models.Harga
	if err := db.Model(h).Order("created_at DESC").First(&h).Error; err != nil {
		log.Printf("error when get data harga from database with error message %s\n", err.Error())
	}

	return h
}
