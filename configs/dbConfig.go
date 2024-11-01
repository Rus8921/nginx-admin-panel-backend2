package configs

import (
	"fmt"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models"
	"log"
)

var Db *gorm.DB

func InitDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
	//log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
	//log.Printf("DB_PASSWORD: %s", os.Getenv("DB_PASSWORD"))
	//log.Printf("DB_NAME: %s", os.Getenv("DB_NAME"))
	//log.Printf("DB_PORT: %s", os.Getenv("DB_PORT"))
	//// Формируем строку подключения из переменных окружения
	//dataForDb := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_NAME"),
	//	os.Getenv("DB_PORT"),
	//)

	dbHost := "localhost"
	dbUser := "your_username"
	dbPassword := "your_password"
	dbName := "your_database_name"
	dbPort := "5432"

	// Form the connection string from hardcoded values
	dataForDb := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)
	log.Printf("Connecting to database with host=%s user=%s dbname=%s port=%s", dbHost, dbUser, dbName, dbPort)

	// Подключаемся к базе данных
	Db, err = gorm.Open(postgres.Open(dataForDb), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	} else {
		log.Printf("successfully connected to database")
	}

	if err := Db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
	Db.Create(&models.User{Username: "admin", HashPassword: "admin"})

}
