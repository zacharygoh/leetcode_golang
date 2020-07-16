package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Config struct {
	DB *gorm.DB
}

var (
	Host     = ""
	Port     = ""
	User     = ""
	Password = ""
	Dbname   = ""
	Url      = ""
)

func Connect() *Config {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		DB: db,
	}
}
