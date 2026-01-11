package main

import (
	"log"
	"net/http"

	"sms-store/db"
	"sms-store/handlers"
	"sms-store/kafka"
)

func main() {
	database, err := db.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}

	go kafka.StartConsumer(database)

	http.HandleFunc(
		"/v1/user/{phoneNumber}/messages",
		handlers.GetHistory(database),
	)

	log.Println("Go SMS Store running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
