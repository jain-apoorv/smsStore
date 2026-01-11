package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"sms-store/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetHistory(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phone := r.PathValue("phoneNumber")

		collection := db.Collection("messages")
		cursor, err := collection.Find(
			context.Background(),
			bson.M{"phoneNumber": phone},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.Background())

		var results []models.SmsEvent
		if err := cursor.All(context.Background(), &results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}
