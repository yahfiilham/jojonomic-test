package app

import (
	"context"
	"encoding/json"
	"log"

	"input-harga-storage-service/internal/config"
	"input-harga-storage-service/internal/models"
)

func ReadMessage() {
	c := config.NewConfig()
	ctx := context.Background()

	for {
		kf, err := c.Kafka.FetchMessage(ctx)
		if err != nil {
			break
		}
		log.Printf("message at topic/partition/offset %v/%v/%v: %s\n", kf.Topic, kf.Partition, kf.Offset, string(kf.Key))

		var m models.Harga
		if err := json.Unmarshal(kf.Value, &m); err != nil {
			log.Printf("unmarshall data error : %s", err.Error())
		}

		if err := c.DB.Model(m).Create(&m).Error; err != nil {
			log.Printf("error insert harga to database with error message : %s \n", err.Error())
		}

		if err := c.Kafka.CommitMessages(ctx, kf); err != nil {
			log.Printf("CommitMessage Failed error : %s", err.Error())
		}

	}
}