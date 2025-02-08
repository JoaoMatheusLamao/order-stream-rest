package main

import (
	"fmt"
	"log"
	"orderstreamrest/internal/config"
	"orderstreamrest/internal/routes"
	"orderstreamrest/internal/utils"
	"orderstreamrest/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println((os.Getenv("ENVIROMENT_EXEC")))

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error creating config: %v", err)
	}
	defer cfg.CloseAll()

	engine := middleware.SetupServer(cfg)

	routes.InitiateRoutes(engine, cfg)

	startServer(engine)
}

func startServer(engine *gin.Engine) {
	certFile, keyFile := utils.GetCertFiles()
	if certFile != "" && keyFile != "" {
		log.Println("Starting server with TLS...")
		if err := engine.RunTLS(":8080", certFile, keyFile); err != nil {
			log.Fatalf("Error starting TLS server: %v", err)
		}
	} else {
		log.Println("Starting server...")
		if err := engine.Run(":8080"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}
	log.Println("Server started on port 8080")
}
