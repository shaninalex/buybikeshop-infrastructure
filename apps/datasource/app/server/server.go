package server

import (
	"buybikeshop/libs/go/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func ProvideGrpcServer(config *config.Config) *grpc.Server {
	creds, err := credentials.NewServerTLSFromFile(
		config.String("certs.server_pem"),
		config.String("certs.server_key_pem"),
	)
	if err != nil {
		panic(err)
	}
	return grpc.NewServer(
		grpc.Creds(creds),
	)
}
