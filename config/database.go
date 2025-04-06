package config

import (
	"gojeksrepo/ent"
	"log"
	"sync"
)

var dbClient *ent.Client
var once sync.Once

func DatabaseInit() {
	once.Do(func() {
		client, err := ent.Open("postgres", "host=localhost port=54320 user=dev dbname=gojeks password=dev sslmode=disable")

		if err != nil {
			log.Fatalf("Database not connected %v", err)
		}
		//if err := client.Schema.Create(context.Background()); err != nil {
		//	log.Fatalf("failed creating schema resources: %v", err)
		//}
		dbClient = client
	})
}
func GetDB() *ent.Client {
	if dbClient == nil {
		log.Fatalf("DB Client not Init, DatabaseInit first.")
	}
	return dbClient
}
