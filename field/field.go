package gotictactoe

import (
	"errors"
)

// The game battle field
type field struct {
	n int      // Square side
	v [][]byte // n*n square values
}

// Creates a new instance of field n*n
func newField(n int) *field {
	v := make([][]byte, n)

	for i := 0; i < n; i++ {
		v[i] = make([]byte, n)
	}

	return &field{n: n, v: v}
}

// Sets v with coordinate c in field
func (f *field) setValue(c Coordinate2, v byte) error {
	if c.x < 0 || c.x >= f.n || c.y < 0 || c.y >= f.n {
		return errors.New("Requested a cell out of the field")
	}

	f.v[c.x][c.y] = v

	return nil
}

// Finds the cells with value v and returns the coordinates
func (f *field) findCellsWithValue(v byte) []Coordinate2D {
	var co []Coordinate2D

	for x := range f.v {
		for y := range f.v[x] {
			if f.v[x][y] == v {
				co = append(co, *NewCoordinate2D(x, y))
			}
		}
	}

	return co
}
