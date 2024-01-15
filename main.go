package main

import (
	"Food-delivery/api/route"
	"Food-delivery/database"
	"Food-delivery/domain"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var err error

func main() {
	database.DB, err = gorm.Open(postgres.Open(database.DbConnect(database.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app := route.SetupRouter()
	database.DB.AutoMigrate(&domain.User{}, &domain.Basket{}, &domain.Product{}, &domain.Promotion{}, &domain.BasketProduct{}, &domain.PromotionProduct{})
	err = app.Listen((":8000"))
	if err != nil {
		panic(err)
	}
}
