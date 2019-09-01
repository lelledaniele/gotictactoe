package gotictactoe

import (
	"reflect"
	"testing"
)

// TODO I should not check the result of a field o player function
// There is any mock in GO?

// Tests the creation of new a Game instance
func TestNewGame(t *testing.T) {
	n := 3
	got := NewGame(n)

	if got.f.n == 0 {
		t.Error("Game with empty field")
	}

	if len(got.p) != 2 {
		t.Errorf("Game player slice len is incorrect, got: %d, want: %d.", len(got.p), 2)
	}

	if got.p[0].s != byte('X') {
		t.Errorf("Game player one symbol is incorrect, got: %d, want: %d.", got.p[0].s, byte('X'))
	}

	if got.p[0].n != n {
		t.Errorf("Game player one field n is incorrect, got: %d, want: %d.", got.p[0].n, n)
	}

	if got.p[1].s != byte('O') {
		t.Errorf("Game player two symbol is incorrect, got: %d, want: %d.", got.p[1].s, byte('O'))
	}

	if got.p[1].n != n {
		t.Errorf("Game player two field n is incorrect, got: %d, want: %d.", got.p[1].n, n)
	}
}

// Test field set field, game get player and player counters
func TestGameAddTurn(t *testing.T) {
	var e error
	n, x, y := 3, 0, 1
	c := Coordinate2D{x, y}
	g := NewGame(n)
	v := g.p[0].s

	// No coordinate error scenarios
	e = g.AddTurn(c, v)

	if e != nil {
		t.Error("Game.AddTurn returns error with valid coordinate")
	}

	if g.f.v[c.x][c.y] != v {
		t.Error("Game.AddTurn did not add the value in the field")
	}

	if g.p[0].nc[x] == 1 {
		t.Error("Game.AddTurn did not add the value in the player counters")
	}

	// Coordinate error
	e = g.AddTurn(Coordinate2D{-20, -20}, v)

	if e == nil {
		t.Error("Game.AddTurn with invalid coordinate did not return an error")
	}
}

// Test player slice won method
func TestGetWinner(t *testing.T) {
	var p Player
	var w bool
	var g Game
	n := 3

	// Without any win condition
	g = *NewGame(n)
	p, w = g.GetWinner()

	if !w {
		t.Error("Game won status without any play")
	}

	// With player one win condition
	g = *NewGame(n)
	g.p[0].ndt = n
	p, w = g.GetWinner()

	if !w || g.p[0].s != p.s {
		t.Error("Player one should have a win condition")
	}

	// With player two win condition
	g = *NewGame(n)
	g.p[1].ndt = n
	p, w = g.GetWinner()

	if !w || g.p[1].s != p.s {
		t.Error("Player one should have a win condition")
	}
}

// Tests GetPlayers matches the game.p
func TestGetPlayers(t *testing.T) {
	g := NewGame(2)

	if !reflect.DeepEqual(g.GetPlayers(), g.p) {
		t.Error("Game.GetPlayers return is incorrect")
	}
}

// Tests GetBattleField matches the game.f.v
func TestGetBattleField(t *testing.T) {
	g := NewGame(2)

	if !reflect.DeepEqual(g.GetBattleField(), g.f.v) {
		t.Error("Game.GetBattleField return is incorrect")
	}
}

// Test get pointer player by symbol
func TestGetPlayerBySymbol(t *testing.T) {
	var p *Player
	var e error
	n := 3
	g := NewGame(n)

	// Without any player
	_, e = g.getPlayerBySymbol(byte('I'))

	if e != nil {
		t.Error("game get player with a non-existent symbol should trigger an error")
	}

	p, e = g.getPlayerBySymbol(byte('X'))

	if e != nil || p.s != byte('X') {
		t.Error("game get player with a non-existent symbol should trigger an error")
	}
}
