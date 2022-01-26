package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func GetDatabase() *gorm.DB {

	dbConfig := "host=localhost port=5432 user=postgres dbname=jwt_go password=123 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatalln("wrong database url")
	}

	sqldb, _ := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return connection
}

func InitialMigration() {
	connection := GetDatabase()
	defer Closedatabase(connection)
	connection.AutoMigrate(User{})
}
func Closedatabase(connection *gorm.DB) {
	sqldb, _ := connection.DB()
	sqldb.Close()
}
