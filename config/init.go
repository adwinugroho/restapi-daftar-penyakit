package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Connection struct {
		DB *gorm.DB
	}
)

var (
	Instance *Connection
	DBUSE    string = os.Getenv("DBUSER")
	DBPASS   string = os.Getenv("DBPASSWORD")
	DBPORT   string = os.Getenv("DBPORT")
)

func Init() {
	sqlDB, err := sql.Open("postgres", "list_of_diseases")
	if err != nil {
		log.Printf("Error while open connection to DB, cause:%+v\n", err)
		return
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:  fmt.Sprintf("user=%s password=%s port=%s", DBUSE, DBPASS, DBPORT),
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Printf("Error while open connection to postgres, cause:%+v\n", err)
		return
	}

	Instance = &Connection{
		DB: gormDB,
	}
}

func GetInstance() *Connection {
	return Instance
}
