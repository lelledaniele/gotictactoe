package gotictactoe

import "testing"

// Tests the creation of new a Coordinate2D instances
func TestNewCoordinate2D(t *testing.T) {
	x, y := 10, 50
	got := NewCoordinate2D(x, y)

	if got.x != x {
		t.Errorf("Coordinatate2D.x is incorrect, got: %d, want: %d.", got.x, x)
	}

	if got.y != y {
		t.Errorf("Coordinatate2D.y is incorrect, got: %d, want: %d.", got.y, y)
	}
}
