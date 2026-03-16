package connector

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"buybikeshop/libs/go/auth"
	"buybikeshop/libs/go/config"
	"log"

	"google.golang.org/grpc"
)

type DatasourceClient struct {
	conn *grpc.ClientConn

	catalogClient pb.CatalogClient
}

func ProvideClient(config *config.Config) *DatasourceClient {
	creds, err := auth.GrpcCredentials(config.String("creds.ca_pem"))
	if err != nil {
		panic(err)
	}

	conn, err := grpc.NewClient(
		config.String("grpc.host"),
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewCatalogClient(conn)

	return &DatasourceClient{
		conn: conn,

		catalogClient: c,
	}
}

func (c *DatasourceClient) Close() {
	c.conn.Close()
}
