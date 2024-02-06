package config

import "testing"

func TestLoadConfig(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Error(err)
	}
}
