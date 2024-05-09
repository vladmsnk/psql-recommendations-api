package clients

import (
	"fmt"
	docker_client "github.com/docker/docker/client"
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

type DockerClient struct {
	Client *docker_client.Client
}

type RedisClient struct {
}

func (cc *CollectorClient) Close() {
	if cc.Conn != nil {
		cc.Conn.Close()
	}
}

func (cc *DockerClient) Close() {
	if cc.Client != nil {
		cc.Client.Close()
	}
}

func (cc *RedisClient) Close() {

}

func NewDockerClient() (*DockerClient, error) {
	client, err := docker_client.NewClientWithOpts(docker_client.FromEnv, docker_client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("docker_client.NewClientWithOpts: %w", err)
	}
	return &DockerClient{Client: client}, nil
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

func NewRedisClient(config config.Redis) (*RedisClient, error) {
	return &RedisClient{}, nil
}
