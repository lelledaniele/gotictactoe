package gotictactoe

import "testing"

// Tests the creation of a new Player instance
func TestNewPlayer(t *testing.T) {
	s, n := byte('X'), 2
	got := newPlayer(s, n)

	if got.s != s {
		t.Errorf("Player symbol is incorrect, got: %d, want: %d.", got.s, s)
	}

	if got.n != n {
		t.Errorf("Player field square side is incorrect, got: %d, want: %d.", got.n, n)
	}

	if len(got.nr) != n {
		t.Errorf("Player row counter len is incorrect, got: %d, want: %d.", len(got.nr), n)
	}

	// Test default value in nr slice
	for i, v := range got.nr {
		if v != 0 {
			t.Errorf("Player column %d counter is incorrect, got: %d, want: %d.", i, v, 0)
		}
	}

	if len(got.nc) != n {
		t.Errorf("Player column counter len is incorrect, got: %d, want: %d.", len(got.nc), n)
	}

	// Test default value in nr slice
	for i, v := range got.nc {
		if v != 0 {
			t.Errorf("Player column %d counter is incorrect, got: %d, want: %d.", i, v, 0)
		}
	}

	if got.ndt != 0 {
		t.Errorf("Player diagonal top counter is incorrect, got: %d, want: %d.", got.ndt, 0)
	}

	if got.ndb != 0 {
		t.Errorf("Player diagonal bottom counter is incorrect, got: %d, want: %d.", got.ndb, 0)
	}
}

// Tests the player symbol getter
func TestGetSymbol(t *testing.T) {
	s, n := byte('X'), 2
	got := newPlayer(s, n)

	if got.GetSymbol() != s {
		t.Errorf("Player symbol getter is incorrect, got: %d, want: %d.", got.GetSymbol(), s)
	}
}

// Tests the player counters
func TestAddTurn(t *testing.T) {
	var c Coordinate2D
	var e error
	var x, y int
	s, n := byte('X'), 3
	p := newPlayer(s, n)

	// Coordinate with x < 0
	x, y = -20, 0
	c = *NewCoordinate2D(x, y)
	e = p.addTurn(c)

	if e == nil {
		t.Error("Add player turn with coordinate.x < 0, but no error occured")
	}

	// Coordinate with x > p.nr
	x, y = len(p.nr)+1, 0
	c = *NewCoordinate2D(x, y)
	e = p.addTurn(c)

	if e == nil {
		t.Errorf("Add player turn with coordinate.x > len(p.nr)(%d), but no error occured", len(p.nr))
	}

	// Coordinate with x == p.nr
	x, y = len(p.nr), 0
	c = *NewCoordinate2D(x, y)
	e = p.addTurn(c)

	if e == nil {
		t.Errorf("Add player turn with coordinate.x == len(p.nr)(%d), but no error occured", len(p.nr))
	}

	// Coordinate with y < 0
	x, y = 0, -20
	c = *NewCoordinate2D(x, y)
	e = p.addTurn(c)

	if e == nil {
		t.Error("Add player turn with coordinate.y < 0, but no error occured")
	}

	// Coordinate with y > p.nc
	x, y = 0, len(p.nc)+1
	c = *NewCoordinate2D(x, y)
	e = p.addTurn(c)

	if e == nil {
		t.Errorf("Add player turn with coordinate.y > len(p.nc)(%d), but no error occured", len(p.nc))
	}

	// Coordinate with y == p.nc
	x, y = 0, len(p.nc)
	c = *NewCoordinate2D(x, y)
	e = p.addTurn(c)

	if e == nil {
		t.Errorf("Add player turn with coordinate.y == len(p.nc)(%d), but no error occured", len(p.nc))
	}

	// Coordinate in diagonal top
	x, y = 0, 0
	c = *NewCoordinate2D(x, y)
	p = newPlayer(s, n)
	e = p.addTurn(c)

	if e != nil {
		t.Errorf("Player.addTurn returns an error with valid coordinate {%d, %d} ", c.x, c.y)
	}

	if p.nr[c.x] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nr counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nr[c.x], 1)
	}

	if p.nc[c.y] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nc counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nc[c.y], 1)
	}

	if p.ndt != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndt counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndt, 1)
	}

	if p.ndb != 0 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndb counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndb, 0)
	}

	// Coordinate not in diagonal top
	x, y = 1, 0
	c = *NewCoordinate2D(x, y)
	p = newPlayer(s, n)
	e = p.addTurn(c)

	if e != nil {
		t.Errorf("Player.addTurn returns an error with valid coordinate {%d, %d} ", c.x, c.y)
	}

	if p.nr[c.x] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nr counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nr[c.x], 1)
	}

	if p.nc[c.y] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nc counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nc[c.y], 1)
	}

	if p.ndt != 0 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndt counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndt, 0)
	}

	if p.ndb != 0 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndb counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndb, 0)
	}

	// Coordinate in diagonal bottom
	x, y = 2, 0
	c = *NewCoordinate2D(x, y)
	p = newPlayer(s, n)
	e = p.addTurn(c)

	if e != nil {
		t.Errorf("Player.addTurn returns an error with valid coordinate {%d, %d} ", c.x, c.y)
	}

	if p.nr[c.x] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nr counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nr[c.x], 1)
	}

	if p.nc[c.y] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nc counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nc[c.y], 1)
	}

	if p.ndt != 0 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndt counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndt, 0)
	}

	if p.ndb != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndb counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndb, 1)
	}

	// Coordinate not in diagonal bottom
	x, y = 2, 1
	c = *NewCoordinate2D(x, y)
	p = newPlayer(s, n)
	e = p.addTurn(c)

	if e != nil {
		t.Errorf("Player.addTurn returns an error with valid coordinate {%d, %d} ", c.x, c.y)
	}

	if p.nr[c.x] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nr counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nr[c.x], 1)
	}

	if p.nc[c.y] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nc counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nc[c.y], 1)
	}

	if p.ndt != 0 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndt counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndt, 0)
	}

	if p.ndb != 0 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndb counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndb, 0)
	}

	// Coordinate not in diagonal top or bottom
	x, y = 0, 1
	c = *NewCoordinate2D(x, y)
	p = newPlayer(s, n)
	e = p.addTurn(c)

	if e != nil {
		t.Errorf("Player.addTurn returns an error with valid coordinate {%d, %d} ", c.x, c.y)
	}

	if p.nr[c.x] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nr counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nr[c.x], 1)
	}

	if p.nc[c.y] != 1 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the nc counter is incorrect, got: %d, want: %d.", c.x, c.y, p.nc[c.y], 1)
	}

	if p.ndt != 0 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndt counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndt, 0)
	}

	if p.ndb != 0 {
		t.Errorf("Add player turn with coordinate {%d, %d} but the ndb counter is incorrect, got: %d, want: %d.", c.x, c.y, p.ndb, 0)
	}
}

// Tests counters if have the win condition
func TestWon(t *testing.T) {
	var p Player
	s, n := byte('X'), 3

	// Lose empty player
	p = *newPlayer(s, n)

	if p.won() {
		t.Error("Empty player has win condition")
	}

	// Win for diagonal top
	p = *newPlayer(s, n)
	p.ndt = n

	if !p.won() {
		t.Error("Player lose with win diagonal top win condition")
	}

	// Win for diagonal bottom
	p = *newPlayer(s, n)
	p.ndb = n

	if !p.won() {
		t.Error("Player lose with win diagonal bottom win condition")
	}

	// Win for i row
	for i := 0; i < n; i++ {
		p = *newPlayer(s, n)
		p.nr[i] = n

		if !p.won() {
			t.Errorf("Player lose with win row[%d] win condition", i)
		}
	}

	// Win for i column
	for i := 0; i < n; i++ {
		p = *newPlayer(s, n)
		p.nc[i] = n

		if !p.won() {
			t.Errorf("Player lose with win column[%d] win condition", i)
		}
	}
}
