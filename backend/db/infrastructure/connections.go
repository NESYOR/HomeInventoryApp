package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)
var database *gorm.DB
func InitConnection(){
	dsn := "host=172.18.0.2 user=admin password=Admin dbname=inventory_app port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Database connection Error")
	}
	database=db
}
func GetConnection() *gorm.DB {
	if database == nil {
		InitConnection()
	}
	return database
}
