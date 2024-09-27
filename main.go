package main

import (
	"GoProject/config"
	_ "GoProject/docs"
	"GoProject/routes"
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

	//config.DB.AutoMigrate(&Models.Guitar{})
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
