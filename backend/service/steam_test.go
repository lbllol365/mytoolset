package service

import (
	"context"
	"github.com/go-resty/resty/v2"
	"testing"
)

var service SteamService

func initAll() {
	if service.Client == nil {
		client := resty.New()
		client.SetProxy("http://127.0.0.1:1081")
		service = SteamService{
			Ctx:    context.TODO(),
			Client: client,
		}
	}
}

func TestGameSearchSuggest(t *testing.T) {
	initAll()
	data := service.GetGameSearchSuggest("cs")
	t.Log(data)
}

func TestGameDetailInfo(t *testing.T) {
	initAll()
	data := service.GetGameDetailInfo(572220)
	t.Log(data)
}

func TestGameWorkShopData(t *testing.T) {
	initAll()
	data := service.getWorkshopDataFromAppid(4000)
	t.Log(data)
}
