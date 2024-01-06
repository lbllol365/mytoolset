package service

import (
	"changeme/backend/types"
	"context"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
	"strings"
)

type SteamService struct {
	ctx context.Context
}

type SteamGameInfo struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	ImageUrl string `json:"imageUrl"`
	Price    string `json:"price"`
}

// getGameSearchSuggest 请求向steam搜索建议接口
func (s SteamService) GetGameSearchSuggest(name string) (resp types.JSResp) {
	urlTemplate := "https://store.steampowered.com/search/suggest?term=%s&f=games&cc=CN&realm=1&l=schinese"
	url := fmt.Sprintf(urlTemplate, name)
	url = strings.Replace(url, " ", "%20", -1)
	fmt.Println(url)
	client := resty.New()
	// TODO 代理可配置
	client.SetProxy("http://127.0.0.1:1081")
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0")
	response, err := client.R().Get(url)
	if err != nil {
		resp.Success = false
		resp.Msg = "steam搜索建议接口请求失败"
		return
	}

	rawContent := string(response.Body())
	parse, err := html.Parse(strings.NewReader(rawContent))
	if err != nil {
		resp.Success = false
		resp.Msg = "解析steam返回的HTML失败"
		return
	}
	query := htmlquery.Find(parse, "//a")
	var dataList []SteamGameInfo
	for _, node := range query {
		gameUrl := htmlquery.SelectAttr(node, "href")
		nameNode := htmlquery.FindOne(node, "//div[@class='match_name']")
		gameName := htmlquery.InnerText(nameNode)
		imageNode := htmlquery.FindOne(node, "//img")
		imageUrl := htmlquery.SelectAttr(imageNode, "src")
		priceNode := htmlquery.FindOne(node, "//div[@class='match_price']")
		price := htmlquery.InnerText(priceNode)
		dataList = append(dataList, SteamGameInfo{
			Name:     gameName,
			ImageUrl: imageUrl,
			Url:      gameUrl,
			Price:    price,
		})
	}
	resp.Success = true
	resp.Data = dataList
	return
}
