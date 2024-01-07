package service

import (
	"changeme/backend/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
	"regexp"
	"strconv"
	"strings"
)

type SteamService struct {
	Ctx    context.Context
	Client *resty.Client
}

type SteamGameInfo struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	ImageUrl string `json:"imageUrl"`
	Price    string `json:"price"`
}

// GetGameSearchSuggest 请求向steam搜索建议接口
func (s SteamService) GetGameSearchSuggest(name string) (resp types.JSResp) {
	urlTemplate := "https://store.steampowered.com/search/suggest?term=%s&f=games&cc=CN&realm=1&l=schinese"
	url := fmt.Sprintf(urlTemplate, name)
	url = strings.Replace(url, " ", "%20", -1)
	fmt.Println(url)
	response, err := s.Client.R().Get(url)
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

// GetGameDetailInfo 根据appid获取游戏详细信息
func (s SteamService) GetGameDetailInfo(appid int) (resp types.JSResp) {
	appidStr := strconv.Itoa(appid)
	detailInfoUrl := "https://store.steampowered.com/api/appdetails/?appids=" + appidStr
	response, err := s.Client.R().Get(detailInfoUrl)
	if err != nil {
		resp.Success = false
		resp.Msg = "请求steam游戏详细接口失败"
		return
	}
	var detailData DetailInfo
	err = json.Unmarshal(response.Body()[len(appidStr)+4:len(response.Body())-1], &detailData)
	if err != nil {
		resp.Success = false
		resp.Msg = "接口返回消息反序列化失败"
		return
	}
	resp.Success = true
	resp.Data = &detailData
	return
}

// getWorkshopDataFromAppid 根据appid获取游戏创意工坊相关信息
func (s SteamService) getWorkshopDataFromAppid(appid int) (workShopInfo WorkShopInfo) {
	urlTemplate := "https://steamcommunity.com/app/%s/workshop?cc=CN&realm=1&l=schinese"
	url := fmt.Sprintf(urlTemplate, strconv.Itoa(appid))
	response, err := s.Client.R().Get(url)
	if err != nil {
		workShopInfo.Success = false
		return
	}
	workShopInfo.Success = true
	// 请求被重定向到商店首页，说明该游戏没有创意工坊
	if response.StatusCode() == 302 {
		workShopInfo.Have = false
		return
	}
	itemNumRe := regexp.MustCompile("<span>查看所有.*</span>")
	itemNum := itemNumRe.FindString(string(response.Body()))
	if itemNum == "" {
		workShopInfo.Success = false
		return
	}
	numRe := regexp.MustCompile("[0-9|,]+")
	num := numRe.FindString(itemNum)
	workShopInfo.Have = true
	workShopInfo.Num = num
	workShopInfo.Link = url
	return
}

type WorkShopInfo struct {
	Success bool   // 请求是否成功
	Have    bool   // 是否有创意工坊
	Num     string // 创意工坊物品数量
	Link    string // 创意工坊链接
}

type DetailInfo struct {
	Success bool `json:"success"`
	Data    struct {
		Type                string `json:"type"`
		Name                string `json:"name"`
		SteamAppid          int    `json:"steam_appid"`
		RequiredAge         int    `json:"required_age"`
		IsFree              bool   `json:"is_free"`
		ControllerSupport   string `json:"controller_support"`
		DetailedDescription string `json:"detailed_description"`
		AboutTheGame        string `json:"about_the_game"`
		ShortDescription    string `json:"short_description"`
		SupportedLanguages  string `json:"supported_languages"`
		HeaderImage         string `json:"header_image"`
		CapsuleImage        string `json:"capsule_image"`
		CapsuleImagev5      string `json:"capsule_imagev5"`
		Website             string `json:"website"`
		PcRequirements      struct {
			Minimum string `json:"minimum"`
		} `json:"pc_requirements"`
		MacRequirements struct {
			Minimum string `json:"minimum"`
		} `json:"mac_requirements"`
		LinuxRequirements struct {
			Minimum string `json:"minimum"`
		} `json:"linux_requirements"`
		Developers    []string `json:"developers"`
		Publishers    []string `json:"publishers"`
		PriceOverview struct {
			Currency         string `json:"currency"`
			Initial          int    `json:"initial"`
			Final            int    `json:"final"`
			DiscountPercent  int    `json:"discount_percent"`
			InitialFormatted string `json:"initial_formatted"`
			FinalFormatted   string `json:"final_formatted"`
		} `json:"price_overview"`
		Packages      []int `json:"packages"`
		PackageGroups []struct {
			Name                    string `json:"name"`
			Title                   string `json:"title"`
			Description             string `json:"description"`
			SelectionText           string `json:"selection_text"`
			SaveText                string `json:"save_text"`
			DisplayType             int    `json:"display_type"`
			IsRecurringSubscription string `json:"is_recurring_subscription"`
			Subs                    []struct {
				Packageid                int    `json:"packageid"`
				PercentSavingsText       string `json:"percent_savings_text"`
				PercentSavings           int    `json:"percent_savings"`
				OptionText               string `json:"option_text"`
				OptionDescription        string `json:"option_description"`
				CanGetFreeLicense        string `json:"can_get_free_license"`
				IsFreeLicense            bool   `json:"is_free_license"`
				PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
			} `json:"subs"`
		} `json:"package_groups"`
		Platforms struct {
			Windows bool `json:"windows"`
			Mac     bool `json:"mac"`
			Linux   bool `json:"linux"`
		} `json:"platforms"`
		Categories []struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
		} `json:"categories"`
		Genres []struct {
			ID          string `json:"id"`
			Description string `json:"description"`
		} `json:"genres"`
		Screenshots []struct {
			ID            int    `json:"id"`
			PathThumbnail string `json:"path_thumbnail"`
			PathFull      string `json:"path_full"`
		} `json:"screenshots"`
		Movies []struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Thumbnail string `json:"thumbnail"`
			Webm      struct {
				Num480 string `json:"480"`
				Max    string `json:"max"`
			} `json:"webm"`
			Mp4 struct {
				Num480 string `json:"480"`
				Max    string `json:"max"`
			} `json:"mp4"`
			Highlight bool `json:"highlight"`
		} `json:"movies"`
		Recommendations struct {
			Total int `json:"total"`
		} `json:"recommendations"`
		Achievements struct {
			Total       int `json:"total"`
			Highlighted []struct {
				Name string `json:"name"`
				Path string `json:"path"`
			} `json:"highlighted"`
		} `json:"achievements"`
		ReleaseDate struct {
			ComingSoon bool   `json:"coming_soon"`
			Date       string `json:"date"`
		} `json:"release_date"`
		SupportInfo struct {
			URL   string `json:"url"`
			Email string `json:"email"`
		} `json:"support_info"`
		Background         string `json:"background"`
		BackgroundRaw      string `json:"background_raw"`
		ContentDescriptors struct {
			Ids   []any `json:"ids"`
			Notes any   `json:"notes"`
		} `json:"content_descriptors"`
	} `json:"data"`
}
