package gotictactoe

import "fmt"

// Coordinate2D x,y
// I use int instead of uint here
// to avoid type conversion during ints comparison, For example with `len()`
type Coordinate2D struct {
	x, y int
}

// NewCoordinate2D creates a new instance of Coordinate2D
func NewCoordinate2D(x, y int) *Coordinate2D {
	return &Coordinate2D{x, y}
}

// String returns the coordinate in {x, y} format
func (c *Coordinate2D) String() string {
	return fmt.Sprintf("{%d, %d}", c.x, c.y)
}
