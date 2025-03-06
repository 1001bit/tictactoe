package room

import (
	"math/rand"
	"time"
)

type Game struct {
	board [3][3]byte
	turn  byte

	rows  [3][2]int
	cols  [3][2]int
	diags [2][2]int
	n     int
}

func NewGame() *Game {
	g := &Game{
		board: [3][3]byte{},
		turn:  ' ',

		rows:  [3][2]int{{0, 0}, {0, 0}, {0, 0}},
		cols:  [3][2]int{{0, 0}, {0, 0}, {0, 0}},
		diags: [2][2]int{{0, 0}, {0, 0}},
		n:     9,
	}
	return g
}

func (g *Game) Start() {
	g.board = [3][3]byte{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}

	g.rows = [3][2]int{{0, 0}, {0, 0}, {0, 0}}
	g.cols = [3][2]int{{0, 0}, {0, 0}, {0, 0}}
	g.diags = [2][2]int{{0, 0}, {0, 0}}
	g.n = 9

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	turn := byte('X')
	if r.Intn(2) == 0 {
		turn = byte('O')
	}
	g.turn = turn
}

func (g *Game) CheckVictory(x, y int) byte {
	for i := 0; i < 2; i++ {
		if g.rows[x][i] == 3 || g.cols[y][i] == 3 || g.diags[0][i] == 3 || g.diags[1][i] == 3 {
			if i == 0 {
				return 'X'
			} else {
				return 'O'
			}
		}
	}

	if g.n == 0 {
		return 'D'
	}
	return ' '
}

func (g *Game) Place(x, y int) bool {
	if x < 0 || x > 2 || y < 0 || y > 2 || g.board[x][y] != ' ' {
		return false
	}

	g.board[x][y] = g.turn
	g.n -= 1

	if g.turn == 'X' {
		g.rows[x][0]++
		g.cols[y][0]++
		if x == y {
			g.diags[0][0]++
		}
		if x == 2-y {
			g.diags[1][0]++
		}

		g.turn = 'O'
	} else {
		g.rows[x][1]++
		g.cols[y][1]++
		if x == y {
			g.diags[0][1]++
		}
		if x == 2-y {
			g.diags[1][1]++
		}

		g.turn = 'X'
	}
	return true
}

func (g *Game) GetTurn() byte {
	return g.turn
}

func (g *Game) GetBoard() [3][3]byte {
	return g.board
}
