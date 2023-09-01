package game_test

import (
	"encoding/json"
	"fmt"
	"lifegame/game"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := game.NewGame(10, 10)
	g.InitLife(50)
	strings, _ := json.Marshal(g)
	fmt.Println(string(strings))
	g.Print()
	fmt.Println("____________________")
	g.NextRound()
	g.Print()
}
