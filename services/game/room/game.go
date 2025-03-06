package room

import (
	"math/rand"
	"time"
)

type Game struct {
	board [3][3]byte
	turn  byte
}

func NewGame() *Game {
	g := &Game{
		board: [3][3]byte{},
		turn:  ' ',
	}
	return g
}

func (g *Game) Start() {
	g.board = [3][3]byte{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	turn := byte('X')
	if r.Intn(2) == 0 {
		turn = byte('O')
	}
	g.turn = turn
}

func (g *Game) Place(x, y int) bool {
	if x < 0 || x > 2 || y < 0 || y > 2 || g.board[x][y] != ' ' {
		return false
	}

	g.board[x][y] = g.turn

	if g.turn == 'X' {
		g.turn = 'O'
	} else {
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
