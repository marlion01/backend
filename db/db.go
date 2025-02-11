package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println("Connecting to DB")
	//fmt.Println(os.Getenv("POSTGRES_USER"))
	//fmt.Println(os.Getenv("POSTGRES_PW"))
	//fmt.Println(os.Getenv("POSTGRES_HOST"))
	//fmt.Println(os.Getenv("POSTGRES_PORT"))
	//fmt.Println(os.Getenv("POSTGRES_DB"))
	shell := os.Getenv("POSTGRES_USER")
	if shell == "" {
		fmt.Println("POSTGRES_USER is empty")
		return nil
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
