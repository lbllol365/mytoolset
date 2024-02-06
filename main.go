package main

import (
	"changeme/backend/config"
	"changeme/backend/db"
	"changeme/backend/service"
	"context"
	"embed"
	"github.com/go-resty/resty/v2"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	ctx := context.TODO()
	client := resty.New()
	errConfig := config.LoadConfig()
	// TODO 日志
	if errConfig != nil {
		panic("加载配置文件失败")
	}
	client.SetProxy(config.Config.ProxyConfig.Host)
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0")
	rssService := &service.RssService{
		Ctx: ctx,
	}
	steamService := &service.SteamService{
		Ctx:    ctx,
		Client: client,
	}
	clientService := &service.DockerService{
		Ctx: ctx,
	}
	clientService.InitClient()
	db.InitDB()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "mytoolset",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			rssService,
			steamService,
			clientService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
