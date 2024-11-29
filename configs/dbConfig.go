package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Admin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/NginxServer"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/SSLcertificat"
	models "gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Site"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Upstreams"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/User"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	if err := Db.AutoMigrate(&User.User{}); err != nil {
		log.Fatalf("failed to migrate user database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
	if err := Db.AutoMigrate(&NginxServer.NginxServer{}); err != nil {
		log.Fatalf("failed to migrate server database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
	if err := Db.AutoMigrate(&models.Site{}); err != nil {
		log.Fatalf("failed to migrate site database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
	if err := Db.AutoMigrate(&Admin.Admin{}); err != nil {
		log.Fatalf("failed to migrate admin database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
	if err := Db.AutoMigrate(&Permission.Permission{}); err != nil {
		log.Fatalf("failed to migrate permission database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
	if err := Db.AutoMigrate(&SSLcertificat.SSL{}); err != nil {
		log.Fatalf("failed to migrate ssl database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
	if err := Db.AutoMigrate(&Upstreams.Upstream{}); err != nil {
		log.Fatalf("failed to migrate upstream database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
	if err := Db.AutoMigrate(&models.Location{}); err != nil {
		log.Fatalf("failed to migrate location database: %v", err)
	} else {
		log.Printf("successfully migrated database")
	}
}
