package gotictactoe

import (
	"fmt"
	"testing"
)

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

func TestCoordinate2D_StringString(t *testing.T) {
	x, y := 10, 50
	expected := fmt.Sprintf("{%d, %d}", x, y)
	c := NewCoordinate2D(x, y)
	got := c.String()

	if got != expected {
		t.Errorf("Coordinatate2D.String is incorrect, got: '%v', want: '%v'.", got, expected)
	}
}
