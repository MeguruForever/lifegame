package config_test

import (
	"lifegame/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	g := config.LoadConfig(100, 100)
	g.Print()
}
