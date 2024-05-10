package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"

	"gopkg.in/yaml.v3"
	"psqlRecommendationsApi/cmd/clients"
	model "psqlRecommendationsApi/internal/model/discovery"
	"strconv"
)

const imageName = "vladmsnk/psql-collector:latest"

type Adapter interface {
	CreateInstance(ctx context.Context, instanceName string, config []byte) (model.CollectorInstance, error)
}

type Implementation struct {
	dockerClient *clients.DockerClient
}

func New(dockerClient *clients.DockerClient) *Implementation {
	return &Implementation{
		dockerClient: dockerClient,
	}
}

func (i *Implementation) CreateInstance(ctx context.Context, instanceName string, config []byte) (model.CollectorInstance, error) {
	var containerId string

	envs, err := GetEnsFromConfig(config)
	if err != nil {
		return model.CollectorInstance{}, fmt.Errorf("GetEnsFromConfig: %w", err)
	}
	_, err = i.dockerClient.Client.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return model.CollectorInstance{}, err
	}
	containers, err := i.dockerClient.Client.ContainerList(ctx, container.ListOptions{
		Filters: filters.NewArgs(filters.Arg("name", instanceName)),
		All:     true,
	})
	if err != nil {
		return model.CollectorInstance{}, fmt.Errorf("dockerClient.Client.ContainerList: %w", err)
	}
	if len(containers) == 0 {
		createResponse, err := i.dockerClient.Client.ContainerCreate(ctx, &container.Config{Env: envs, Cmd: []string{"/app"}, Image: imageName}, &container.HostConfig{}, &network.NetworkingConfig{map[string]*network.EndpointSettings{"shared_network": {}}}, nil, instanceName)
		if err != nil {
			return model.CollectorInstance{}, err
		}
		containerId = createResponse.ID
	} else {
		containerId = containers[0].ID
	}
	err = i.dockerClient.Client.ContainerStart(ctx, containerId, container.StartOptions{})
	if err != nil {
		return model.CollectorInstance{}, err
	}

	return model.CollectorInstance{Id: containerId, Host: instanceName}, nil
}

func GetEnsFromConfig(config []byte) ([]string, error) {
	p := Postgres{}
	err := yaml.Unmarshal(config, &p)
	if err != nil {
		return nil, err
	}
	envs := []string{
		"PG_USER=" + p.User,
		"PG_PASSWORD=" + p.Password,
		"PG_DATABASE=" + p.Database,
		"PG_SSLMODE=" + p.SSLMode,
		"PG_HOST=" + p.Host,
		"PG_PORT=" + strconv.Itoa(p.Port),
	}
	return envs, nil
}
