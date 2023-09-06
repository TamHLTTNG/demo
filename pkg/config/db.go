package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
)

func NewDB() (*gorm.DB, error) {
	var (
		DB_Host   = C.DB.Host
		DB_Port   = C.DB.Port
		DB_DBName = C.DB.Name
		DB_User   = C.DB.User
		DB_Passwd = C.DB.Password
		DB_Driver = C.DB.Driver
	)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	DBMS := "postgres"
	postgresConfig := &postgres.Config{
		DriverName: DB_Driver,
		DSN: fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			DB_Host, DB_Port, DB_User, DB_Passwd, DB_DBName),
	}

	print(postgresConfig.DSN)

	db, err := gorm.Open(DBMS, postgresConfig.DSN)

	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}
