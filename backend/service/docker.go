package service

import (
	"changeme/backend/types"
	"context"
	types2 "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerService struct {
	Ctx    context.Context
	Client *client.Client
}

func (s *DockerService) InitClient() (resp types.JSResp) {
	apiClient, err := client.NewClientWithOpts(client.WithHost("http://192.168.0.102:2376"), client.WithVersion("1.41"))
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
