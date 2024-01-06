package service

import (
	"changeme/backend/types"
	"context"
	"github.com/mmcdole/gofeed"
)

type RssService struct {
	ctx context.Context
}

func (r *RssService) PullData() (resp types.JSResp) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://game.ali213.net/forum.php?mod=rss&fid=77")
	if err != nil {
		resp.Success = false
		resp.Msg = "Feed Parse Error"
		return
	}
	resp.Data = &feed.Items
	resp.Success = true
	return
}
