package gotictactoe

import "errors"

// TODO - Find a (nice) way to remove any reference about the game field in the player properties

// Player for the game
type Player struct {
	s        byte  // Symbol
	n        int   // The battle's field squre side
	ndt, ndb int   // Counter for player turns in diagolan top and bottom
	nr, nc   []int // Counter for player turns in rows and cells
}

// Creates a new instance of Player
func newPlayer(s byte, n int) *Player {
	return &Player{
		s:  s,
		n:  n,
		nr: make([]int, n),
		nc: make([]int, n),
	}
}

// GetSymbol exposes player symbol
func (p *Player) GetSymbol() byte {
	return p.s
}

// Increments the stats ndt/ndb/nr/nc for the player
func (p *Player) addTurn(c Coordinate2D) error {
	if c.x < 0 || c.x >= len(p.nr) || c.y < 0 || c.y >= len(p.nc) {
		return errors.New("Coordinate out of the player counters")
	}

	p.nr[c.x]++
	p.nc[c.y]++

	// Diagonal top. Ex. {0,0} {1,1} {2,2}
	if c.x == c.y {
		p.ndt++
	}

	// Diagonal bottom. Ex. {2,0} {1,1} {0,2}
	if c.x+c.y == p.n-1 {
		p.ndb++
	}

	return nil
}

// Checks if the player counters have the win condition
func (p *Player) won() bool {
	// The player won if the have p.n play in one of the diagonal
	if p.ndt == p.n || p.ndb == p.n {
		return true
	}

	// The player won if the have p.n play in one of the row/column
	for i := 0; i < p.n; i++ {
		if p.nr[i] == p.n || p.nc[i] == p.n {
			return true
		}
	}

	return false
}
