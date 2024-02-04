package service

import (
	"context"
	"testing"
)

var dockerService DockerService = DockerService{
	Ctx: context.Background(),
}

func initDockerService() {
	dockerService.InitClient()
}

func TestDockerService_PingClient(t *testing.T) {
	initDockerService()
	resp := dockerService.PingClient()
	t.Log(resp)
}

func TestDockerService_ListImage(t *testing.T) {
	initDockerService()
	resp := dockerService.ListImage()
	t.Log(resp)
}
