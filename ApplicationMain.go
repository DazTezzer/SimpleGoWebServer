// @title           Bebeziansback
// @version         1.0
// @description     This is a Bebeziansback
// @host            localhost:8081
// @BasePath        /
package main

import (
	_ "bebeziansback/docs"
	"bebeziansback/server/config"
	"bebeziansback/server/routes"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := InitDB(); err != nil {
		log.Fatalf("ERROR: Ошибка соединения с базой данных: %v", err)
	}

	if err := config.RunMigrations(config.BuildDBConfig()); err != nil {
		log.Fatal(err)
	}

	r := routes.SetupRouter()
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("ERROR: Не удалось запустить сервер: %v", err)
	}
}

func InitDB() error {
	var err error
	config.DB, err = gorm.Open(postgres.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("INFO: Успешно соединено с базой данных!")
	return nil
}
