package connector

import (
	"buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"buybikeshop/gen/grpc-buybikeshop-go/partners"
	"buybikeshop/libs/go/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DatasourceClient struct {
	conn *grpc.ClientConn

	CatalogClient  catalog.CatalogClient
	PartnersClient partners.PartnersClient
}

func ProvideDatasourceClient(config *config.Config) *DatasourceClient {
	conn, err := grpc.NewClient(config.String("grpc.host"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &DatasourceClient{
		conn: conn,

		CatalogClient:  catalog.NewCatalogClient(conn),
		PartnersClient: partners.NewPartnersClient(conn),
	}
}

func (c *DatasourceClient) Close() {
	c.conn.Close()
}
