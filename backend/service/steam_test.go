package service

import (
	"context"
	"testing"
)

func TestGameSearchSuggest(t *testing.T) {
	ctx := context.Background()
	steamService := SteamService{ctx: ctx}
	data := steamService.GetGameSearchSuggest("cs")
	t.Log(data)
}
