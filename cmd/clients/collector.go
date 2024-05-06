package clients

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"psqlRecommendationsApi/internal/config"
	pb_collector "psqlRecommendationsApi/pkg/collector"
	pb_redommendation "psqlRecommendationsApi/pkg/recommendations_api"
)

type CollectorClient struct {
	Client pb_collector.CollectorClient
	Conn   *grpc.ClientConn
}

type RecommendationApiClient struct {
	Client pb_redommendation.RecommendationsAPIClient
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

	client := pb_collector.NewCollectorClient(conn)

	return &CollectorClient{Client: client, Conn: conn}, err
}

func NewRecommendationApiClient(config config.RecommendationApi) (*RecommendationApiClient, error) {

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	client := pb_redommendation.NewRecommendationsAPIClient(conn)

	return &RecommendationApiClient{Client: client, Conn: conn}, err
}
