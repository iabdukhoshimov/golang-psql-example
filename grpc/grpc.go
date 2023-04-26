package grpc

import (
	"golang-psql/config"
	"golang-psql/genproto/book_service"
	"golang-psql/grpc/client"
	"golang-psql/grpc/service"
	"golang-psql/pkg/logger"
	"golang-psql/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	book_service.RegisterBookServiceServer(grpcServer, service.NewBookService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
