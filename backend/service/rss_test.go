package service

import (
	"context"
	"encoding/json"
	"testing"
)

func TestPullData(t *testing.T) {
	ctx := context.Background()
	rssService := RssService{Ctx: ctx}
	data := rssService.PullData()
	marshal, _ := json.Marshal(data)
	t.Log(string(marshal))
}
