package app

import (
	"context"
	"encoding/json"
	"log"

	"buyback-storage-service/configs"
	"buyback-storage-service/internal/models"

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

		m := new(models.BuybackData)
		if err := json.Unmarshal(kf.Value, &m); err != nil {
			log.Printf("unmarshall data error : %s", err.Error())
		}

		if err := c.DB.Model(m).Create(m).Error; err != nil {
			log.Printf("error when insert topup to database with error message %s\n", err.Error())
		}

		go saveToRekening(c.DB, m)
		go saveToTransaksi(c.DB, m)

		if err := c.Kafka.CommitMessages(ctx, kf); err != nil {
			log.Printf("CommitMessage Failed error : %s", err.Error())
		}
	}
}

func saveToRekening(db *gorm.DB, data *models.BuybackData) {
	rek := getRekening(db, data.Norek)

	saldo := rek.Saldo - data.Gram
	rek.Saldo = saldo

	if err := db.Model(rek).Where("reff_id = ?", rek.ReffID).Updates(&rek).Error; err != nil {
		log.Printf("error when insert rekening to database with error message %s\n", err.Error())
	}
}

func saveToTransaksi(db *gorm.DB, data *models.BuybackData) {
	rek := getRekening(db, data.Norek)

	id, _ := shortid.Generate()
	m := &models.Transaksi{
		ReffID: id,
		TransaksiData: models.TransaksiData{
			Type:         "buyback",
			Gram:         data.Gram,
			Saldo:        rek.Saldo - data.Gram,
			HargaTopup:   data.HargaTopup,
			HargaBuyback: data.HargaBuyback,
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
