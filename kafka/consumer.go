package kafka

import (
	"context"
	"encoding/json"
	"log"

	"sms-store/models"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartConsumer(db *mongo.Database) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "sms-events",
		GroupID:"sms-store-group",
	})

	collection := db.Collection("messages")

	log.Println("Kafka consumer started...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		var event models.SmsEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Println("JSON unmarshal error:", err)
			continue
		}

		_, err = collection.InsertOne(context.Background(), event)
		if err != nil {
			log.Println("Mongo insert error:", err)
			continue
		}

		log.Printf("Stored SMS event for %s\n", event.PhoneNumber)
	}
}
