package gotictactoe

import (
	"errors"
)

// Game for TicTacToe
type Game struct {
	f field
	p []Player
}

// NewGame match
func NewGame(n int) *Game {
	f := *newField(n)

	p1, p2 := *newPlayer('X', n), *newPlayer('O', n)
	p := []Player{p1, p2}

	return &Game{f: f, p: p}
}

// AddTurn sets value in field and adds turn for the player
func (g *Game) AddTurn(c Coordinate2D, s byte) error {
	e := g.f.setValue(c, s)

	if e != nil {
		return e
	}

	p, e := g.getPlayerBySymbol(s)

	if e != nil {
		return e
	}

	e = p.addTurn(c)

	return e
}

// GetWinner checks if any player won
func (g *Game) GetWinner() (Player, bool) {
	for i := range g.p {
		if g.p[i].won() {
			return g.p[i], true
		}
	}

	return Player{}, false
}

// Gets player by symbol from g.p slice
func (g *Game) getPlayerBySymbol(s byte) (*Player, error) {
	p := new(Player)

	for i := range g.p {
		if g.p[i].s == s {
			p = &g.p[i]
		}
	}

	if p == nil {
		e := errors.New("Player does not exist in the game")

		return p, e
	}

	return p, nil
}
