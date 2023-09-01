package config

import (
	"encoding/json"
	"fmt"
	"io"
	"lifegame/game"
	"log"
	"os"
)

// 配置文件
type Config struct {
	Row               int            `json:"row"`
	Col               int            `json:"col"`
	IfUseAlivePercent bool           `json:"ifUseAlivePercent"`
	AlivePercent      int            `json:"alivePercent"`
	Grid              [][]game.State `json:"grid"`
}

// 读取配置文件
func LoadConfig(Row, Col int) *game.Game {
	f, err := os.Open("/config.json")
	if err != nil {
		fmt.Println("config.json not found, use default config")
		g := game.NewGame(Row, Col)
		g.InitLife(30)
		return g
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	g := game.NewGame(Row, Col)
	if cfg.IfUseAlivePercent {
		g.InitLife(cfg.AlivePercent)
	} else {
		g.Grid = cfg.Grid
	}

	return g
}
