package app

import (
	"cach-plan.go/app/model"
	"cach-plan.go/config"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Paris",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB, err = model.DBMigrate(db)
	if err != nil {
		log.Fatal("Could create schema database")
	}
	a.Router = mux.NewRouter()
	a.setRouters()
}
