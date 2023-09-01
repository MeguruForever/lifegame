package game

import "math/rand"

// 生命状态
type State struct {
	NowState  bool
	NextState bool
}

// 游戏主体对象
type Game struct {
	Row  int       `json:"row"`
	Col  int       `json:"col"`
	Grid [][]State `json:"grid"`
}

//导出为json

// 创建游戏对象
func NewGame(row, col int) *Game {
	game := Game{
		Row: row,
		Col: col,
	}
	game.Grid = Grid(row, col)
	return &game
}

// 生成游戏网格
func Grid(row, col int) [][]State {
	g := make([][]State, row)
	for i := 0; i < row; i++ {
		g[i] = make([]State, col)
	}
	return g
}

// 初始化随机生命
func (game *Game) InitLife(AlivePercent int) {
	for i := 0; i < game.Row; i++ {
		for j := 0; j < game.Col; j++ {
			if rand.Intn(100) < AlivePercent {
				game.Grid[i][j] = State{NowState: true, NextState: true}
			} else {
				game.Grid[i][j] = State{NowState: false, NextState: false}
			}
		}
	}
}

// 判断边界
func (game *Game) MapX(x int) int {
	if x >= game.Row {
		x = 0
	} else if x < 0 {
		x = game.Row - 1
	}
	return x
}
func (game *Game) MapY(y int) int {
	if y >= game.Col {
		y = 0
	} else if y < 0 {
		y = game.Col - 1
	}
	return y
}

// 计算周围生命数量
func (game *Game) CountLife(x, y int) int {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if game.Grid[game.MapX(x+i)][game.MapY(y+j)].NowState {
				count++
			}
		}
	}
	return count
}

// 计算下一步生命状态
func (game *Game) NextState(x, y int) bool {
	count := game.CountLife(x, y)
	if (count == 3 || count == 2) && game.Grid[x][y].NowState {
		return true
	} else if (count < 2 || count > 3) && game.Grid[x][y].NowState {
		return false
	} else if (count == 3) && !game.Grid[x][y].NowState {
		return true
	} else {
		return false
	}
}

// 计算下一步所有生命状态
func (game *Game) CalState() {
	for i := 0; i < game.Row; i++ {
		for j := 0; j < game.Col; j++ {
			game.Grid[i][j].NextState = game.NextState(i, j)
		}
	}
}

// 更新生命状态
func (game *Game) Update() {
	for i := 0; i < game.Row; i++ {
		for j := 0; j < game.Col; j++ {
			game.Grid[i][j].NowState = game.Grid[i][j].NextState
		}
	}
}

// 判断生命是否存活
func (game *Game) Isalive(x, y int) bool {
	return game.Grid[x][y].NowState
}

// 下一回合
func (game *Game) NextRound() {
	game.CalState()
	game.Update()
}

// 打印状态
func (game *Game) Print() {
	for i := 0; i < game.Row; i++ {
		for j := 0; j < game.Col; j++ {
			if game.Isalive(i, j) {
				print("* ")
			} else {
				print("x ")
			}
		}
		println()
	}
}
