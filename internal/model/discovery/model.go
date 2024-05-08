package discovery

import "errors"

var (
	ErrInstanceAlreadyExists = errors.New("instance already exists")
)

type InstanceStatus int64

const (
	InstanceStatusUnspecified InstanceStatus = 0
	InstanceStatusNew         InstanceStatus = 1
	InstanceStatusRunning     InstanceStatus = 2
)

type CollectorInstance struct {
	Id     int64
	Name   string
	Host   string
	Port   int
	Status InstanceStatus
}
