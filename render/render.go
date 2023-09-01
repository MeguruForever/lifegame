package render

import (
	"lifegame/config"
	"lifegame/game"

	"github.com/gdamore/tcell"
)

// 渲染器
type Render struct {
	Game   *game.Game
	Screen tcell.Screen
}

// 创建渲染器
func NewRender(screen tcell.Screen) *Render {
	w, h := screen.Size()
	game := config.LoadConfig(h, w/2)
	return &Render{
		Game:   game,
		Screen: screen,
	}
}

// 绘制
func (r *Render) Draw() {
	r.Screen.Clear()
	for i := 0; i < r.Game.Row; i++ {
		for j := 0; j < r.Game.Col; j++ {
			if r.Game.Grid[i][j].NowState {
				r.Screen.SetCell(j*2, i, tcell.StyleDefault.Background(tcell.ColorWhite), ' ')
				r.Screen.SetCell(j*2+1, i, tcell.StyleDefault.Background(tcell.ColorWhite), ' ')
			} else {
				r.Screen.SetCell(j*2, i, tcell.StyleDefault.Background(tcell.ColorBlack), ' ')
				r.Screen.SetCell(j*2+1, i, tcell.StyleDefault.Background(tcell.ColorBlack), ' ')
			}
		}
	}
	r.Game.NextRound()
	r.Screen.Show()
}
