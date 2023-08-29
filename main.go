package main

import (
	"bwa_startup/handler"
	"bwa_startup/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initializeEnvironment() error {
	environment := os.Getenv("ENV")
	err := godotenv.Load(".env." + environment)
	return err
}

func initializeDatabase() (*gorm.DB, error) {
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	err := initializeEnvironment()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	db, err := initializeDatabase()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	router.Run()

}
