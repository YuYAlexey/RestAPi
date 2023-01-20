package connect

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=todo password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Todoinfo{})

	return db
}
