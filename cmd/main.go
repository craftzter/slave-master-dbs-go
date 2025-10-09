package main

import (
	"log"
	"os"

	"learn-api/internal/delivery/http/handler"
	"learn-api/internal/infrastructure/psql"
	"learn-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Env variable failed to load")
	}

	log.Println("Env loaded successfully")

	// Create uploads directory
	if err := os.MkdirAll("uploads", 0755); err != nil {
		log.Printf("Failed to create uploads directory: %v", err)
	}

	masterDB, err := psql.InitMaster()
	if err != nil {
		panic(err)
	}

	slaveDB, err := psql.InitSlave()
	if err != nil {
		panic(err)
	}

	if masterDB != nil && slaveDB != nil {
		log.Println("PGSQL master and slave connected")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET env variable is required")
	}

	userRepo := psql.NewUserRepoPG(masterDB, slaveDB)
	authUsecase := usecase.NewAuthUsecase(userRepo, jwtSecret)

	profileRepo := psql.NewProfileRepoPG(masterDB, slaveDB)
	profileUsecase := usecase.NewProfileUsecase(profileRepo)

	postRepo := psql.NewPostRepoPG(masterDB, slaveDB)
	postUsecase := usecase.NewPostUsecase(postRepo)

	r := gin.Default()

	handler.NewAuthHandler(authUsecase).RegisterRoutes(r)
	handler.NewProfileHandler(profileUsecase).RegisterRoutes(r)
	handler.NewPostHandler(postUsecase).RegisterRoutes(r)

	// Serve static files for uploads
	r.Static("/uploads", "./uploads")

	log.Println("server running at: 8080")
	r.Run(":8080")
}
