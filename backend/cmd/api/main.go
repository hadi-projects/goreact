package main

import (
	"log"

	"github.com/hadi-projects/go-react-starter/config"
	"github.com/hadi-projects/go-react-starter/internal/router"
	"github.com/hadi-projects/go-react-starter/pkg/database"
	"github.com/hadi-projects/go-react-starter/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	logger.InitLogger(&cfg)

	_, err := database.NewMySQLConnection(&cfg)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	router := router.NewRouter(&cfg)
	router.SetupRouter()
	router.Run()
}
