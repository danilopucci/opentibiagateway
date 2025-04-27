package main

import (
	"fmt"
	"log"
	"net"

	gatewaypb "github.com/danilopucci/opentibiagateway/internal/protogen/v1"
	"github.com/danilopucci/opentibiagateway/internal/provider/mysql"
	"github.com/danilopucci/opentibiagateway/internal/service"
	gatewayGrpcServer "github.com/danilopucci/opentibiagateway/internal/transport/grpc"
	"google.golang.org/grpc"

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

	// Setup gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	grpcServer := gatewayGrpcServer.NewGrpcServer(playerService)

	gatewaypb.RegisterPlayerServiceServer(s, grpcServer)

	fmt.Println("gRPC server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
