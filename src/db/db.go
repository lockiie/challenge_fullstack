package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	//driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

//Pool of connections
var Pool *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())
	}

	// db, err := sql.Open("mysql", os.Getenv("USER_DB")+":"+os.Getenv("PASSWORD_DB")+
	// 	"@tcp("+os.Getenv("HOST_DB")+":"+os.Getenv("PORT_DB")+")/"+os.Getenv("SCHEMA_DB")+"?parseTime=true")
	db, err := sql.Open("mysql", os.Getenv("USER_DB")+":"+os.Getenv("PASSWORD_DB")+
		"@tcp(db)/"+os.Getenv("SCHEMA_DB")+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	Pool = db
	log.Println("Database started")
}
