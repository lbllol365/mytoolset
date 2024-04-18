package service

import (
	"changeme/backend/config"
	"changeme/backend/types"
	"context"
	types2 "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerService struct {
	Ctx    context.Context
	Client *client.Client
}

func (s *DockerService) InitClient() (resp types.JSResp) {
	apiClient, err := client.NewClientWithOpts(client.WithHost(config.Config.DockerConfig.Host), client.WithVersion("1.41"))
	if err != nil {
		resp.Success = false
		resp.Msg = "初始化Docker客户端失败"
	}
	s.Client = apiClient
	resp.Success = true
	return
}

func (s *DockerService) CloseClient() (resp types.JSResp) {
	_ = s.Client.Close()
	resp.Success = true
	return
}

func (s *DockerService) PingClient() (resp types.JSResp) {
	_, err := s.Client.Ping(s.Ctx)
	if err != nil {
		resp.Success = false
		resp.Msg = "Ping客户端失败"
		return
	}
	resp.Success = true
	return
}

// ListImage 获取镜像列表
func (s *DockerService) ListImage() (resp types.JSResp) {
	imageList, err := s.Client.ImageList(s.Ctx, types2.ImageListOptions{
		All: true,
	})
	if err != nil {
		resp.Success = false
		resp.Msg = "获取镜像列表失败"
		return
	}
	resp.Success = true
	resp.Data = &imageList
	return
}

// ListContainer 获取容器列表
func (s *DockerService) ListContainer() (resp types.JSResp) {
	containerList, err := s.Client.ContainerList(s.Ctx, container.ListOptions{
		All: true,
	})
	if err != nil {
		resp.Success = false
		resp.Msg = "获取容器列表失败"
		return
	}
	resp.Success = true
	resp.Data = &containerList
	return
}

// StartContainer 启动指定容器
func (s *DockerService) StartContainer(containerID string) (resp types.JSResp) {
	err := s.Client.ContainerStart(s.Ctx, containerID, container.StartOptions{})
	if err != nil {
		resp.Success = false
		resp.Msg = "启动指定容器失败"
		return
	}
	resp.Success = true
	return
}

// StopContainer 停止指定容器
func (s *DockerService) StopContainer(containerID string, waitTimeout int) (resp types.JSResp) {
	err := s.Client.ContainerStop(s.Ctx, containerID, container.StopOptions{
		Timeout: &waitTimeout,
	})
	if err != nil {
		resp.Success = false
		resp.Msg = "停止指定容器失败"
		return
	}
	resp.Success = true
	return
}
