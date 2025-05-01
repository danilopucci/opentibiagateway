package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	internalLogger "github.com/danilopucci/opentibiagateway/internal/pkg/logger"
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
// - adicionar testes unitarios

const dotEnvFileNamePath = "./../../.env"

func main() {

	logger := internalLogger.NewLoggerBuilder().
		WithTextColorOutput().
		WithLogLevel(slog.LevelDebug).
		Build()

	err := godotenv.Load(dotEnvFileNamePath)
	if err != nil {
		logger.Errorw("error loading .env file", err,
			slog.String("path", dotEnvFileNamePath),
		)
		os.Exit(1)
	}

	dsn := mysql.GenerateDsnFromEnv()

	mysqlDatabase, err := mysql.NewMySqlDatabase(dsn)
	if err != nil {
		logger.Errorw("error connecting to mysql database", err)
		os.Exit(1)
	}

	playerRepository := mysql.NewMySQLPlayerRepository(mysqlDatabase)
	playerService := service.NewPlayerService(playerRepository)

	s := grpc.NewServer()
	grpcServer := gatewayGrpcServer.NewGrpcServer(playerService)
	gatewaypb.RegisterPlayerServiceServer(s, grpcServer)

	// Setup gRPC server
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Errorw("error listening on tcp port", err, slog.Int("port", port))
		os.Exit(1)
	}

	logger.Info("grpc server started", slog.Int("port", port))

	if err := s.Serve(lis); err != nil {
		logger.Errorw("error serving grpc server", err, slog.Int("port", port))
		os.Exit(1)
	}

}
