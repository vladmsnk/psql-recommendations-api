package clients

import (
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"psqlRecommendationsApi/internal/config"
	pb "psqlRecommendationsApi/pkg/collector"
)

type CollectorClient struct {
	Client pb.CollectorClient
	Conn   *grpc.ClientConn
}

func (cc *CollectorClient) Close() {
	if cc.Conn != nil {
		cc.Conn.Close()
	}
}

func NewCollectorClient(config config.Collector) (*CollectorClient, error) {

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	client := pb.NewCollectorClient(conn)

	return &CollectorClient{Client: client, Conn: conn}, err
}
