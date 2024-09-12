package main

import (
	"be-skillacademy-final/api"
	"be-skillacademy-final/db"
	"be-skillacademy-final/model"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
   	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error conver port")
	}

	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         dbHost,
		Username:     dbUser,
		Password:     dbPassword,
		DatabaseName: dbName,
		Port:			port,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate() //&model.User{} as parameter

	// userRepo := repo.NewUserRepo(conn)

	// userService := service.NewUserService(userRepo)
	

	mainAPI := api.NewAPI() //userService as parameter
	mainAPI.Start()
}
