package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var Config *ToolConfig

type ToolConfig struct {
	DockerConfig   *DockerConfig
	DatabaseConfig *DatabaseConfig
	ProxyConfig    *ProxyConfig
}

type DockerConfig struct {
	Host string
}

type DatabaseConfig struct {
	Location string
}

type ProxyConfig struct {
	Host string
}

func LoadConfig() error {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		return fmt.Errorf("加载ini配置文件失败 %v", err)
	}
	Config = &ToolConfig{}
	dockerConfig := &DockerConfig{
		Host: cfg.Section("docker").Key("host").String(),
	}
	Config.DockerConfig = dockerConfig
	databaseConfig := &DatabaseConfig{
		Location: cfg.Section("database").Key("location").String(),
	}
	Config.DatabaseConfig = databaseConfig
	proxyConfig := &ProxyConfig{
		Host: cfg.Section("proxy").Key("host").String(),
	}
	Config.ProxyConfig = proxyConfig
	return nil
}
