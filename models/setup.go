package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error importing .env file")
	}
	DB = Connect()
	log.Println(DB.Stats())
	Migrate()
	// DB.Close()
}

func Connect() *sql.DB {
	connectionStr := os.Getenv("POSTGRES_DSN")
	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	return conn
}


func Migrate() {
	DB.Query(StorageSQL)
	log.Println("Created (if not exists) table storages")
	DB.Query(ItemSQL)
	log.Println("Created (if not exists) table items")
	DB.Query(ClientSQL)
	log.Println("Created (if not exists) table clients")
	DB.Query(CourierSQL)
	log.Println("Created (if not exists) table couriers")
	DB.Query(OrderSQL)
	log.Println("Created (if not exists) table orders")
}