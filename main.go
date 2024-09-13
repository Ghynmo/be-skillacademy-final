package main

import (
	"be-skillacademy-final/api"
	"be-skillacademy-final/db"
	"be-skillacademy-final/model"
	repo "be-skillacademy-final/repository"
	"be-skillacademy-final/service"
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

   	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error convert port")
	}

	db := db.NewDB()
	
	dbCredential := model.Credential{
		Host:         os.Getenv("DB_HOST"),
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port:		  port,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{}, &model.Session{})

	userRepo := repo.NewUserRepo(conn)
	sessionRepo := repo.NewSessionRepo(conn)
	
	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	

	mainAPI := api.NewAPI(userService, sessionService)
	mainAPI.Start()
}
