package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danilopucci/opentibiagateway/internal/provider/mysql"
	"github.com/danilopucci/opentibiagateway/internal/service"

	"github.com/joho/godotenv"
)

// TODO:
// - adicionar o config.yaml - carrega as variaveis de config
// - adicionar um fluxo completo do get player by ID, com GRPC e http server
// - adicionar um logger descente
// - adicionar testes unitarios

const dotEnvFileNamePath = "./../../.env"

func main() {

	err := godotenv.Load(dotEnvFileNamePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := mysql.GenerateDsnFromEnv()

	mysqlDatabase, err := mysql.NewMySqlDatabase(dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	playerRepository := mysql.NewMySQLPlayerRepository(mysqlDatabase)
	playerService := service.NewPlayerService(playerRepository)

	player, err := playerService.GetPlayerByID(context.Background(), 1)
	if err != nil {
		log.Fatalf("Error fetching player: %v", err)
	}

	if player == nil {
		log.Fatalf("player not found\n")
	}

	fmt.Println("Hello, World!")
}
