package docker

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"

	"gopkg.in/yaml.v3"
	"io"
	"os"
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
	p := Postgres{}
	err := yaml.Unmarshal(config, &p)
	if err != nil {
		return model.CollectorInstance{}, err
	}
	envs := []string{
		"PG_USER=" + p.User,
		"PG_PASSWORD=" + p.Password,
		"PG_DATABASE=" + p.Database,
		"PG_SSLMODE=" + p.SSLMode,
		"PG_HOST=" + p.Host,
		"PG_PORT=" + strconv.Itoa(p.Port),
	}

	pullOut, err := i.dockerClient.Client.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return model.CollectorInstance{}, err
	}
	io.Copy(os.Stdout, pullOut)

	containers, err := i.dockerClient.Client.ContainerList(ctx, container.ListOptions{
		Filters: filters.NewArgs(filters.Arg("name", instanceName)),
		All:     true, // Include stopped containers
	})
	if err != nil {
		panic(err)
	}

	var containerId string

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

	return model.CollectorInstance{Id: containerId}, nil
}
