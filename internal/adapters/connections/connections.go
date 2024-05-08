package connections

import (
	"context"
	"fmt"
	"psqlRecommendationsApi/cmd/clients"
	"psqlRecommendationsApi/internal/config"
	"sync"
)

var (
	ErrConnectionNotFound   = fmt.Errorf("Connection not found")
	ErrConnectionAlreadySet = fmt.Errorf("Connection already set")
)

type ConnectionProvider interface {
	GetConnection(ctx context.Context, instanceName string) (*clients.CollectorClient, error)
	SetConnection(ctx context.Context, instanceName, host string, port int64) error
}

type Implementation struct {
	storage map[string]*clients.CollectorClient
	mu      sync.Mutex
}

func (i *Implementation) GetConnection(_ context.Context, instanceName string) (*clients.CollectorClient, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	connection, ok := i.storage[instanceName]
	if !ok {
		return nil, ErrConnectionNotFound
	}

	return connection, nil
}

func (i *Implementation) SetConnection(_ context.Context, instanceName, host string, port int) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, ok := i.storage[instanceName]
	if ok {
		return ErrConnectionAlreadySet
	}

	conn, err := clients.NewCollectorClient(config.Collector{Host: host, Port: port})
	if err != nil {
		return fmt.Errorf(" clients.NewCollectorClient: %w", err)
	}

	i.storage[instanceName] = conn
	return nil
}
