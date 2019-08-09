package gotictactoe

import "testing"

// Tests the creation of new field instances
func TestNewField(t *testing.T) {
	n := 3
	// Field 3x3
	got := newField(n)

	if got.n != n {
		t.Errorf("field.n is incorrect, got: %d, want: %d.", got.n, n)
	}

	if len(got.v) != n {
		t.Errorf("field.v has incorrect number of rows, got: %d, want: %d.", len(got.v), n)
	}

	for i := range got.v {
		if len(got.v[i]) != n {
			t.Errorf("field.v[%d] has incorrect number of columns, got: %d, want: %d.", i, len(got.v[i]), n)
		}

		for y := range got.v[i] {
			// Every cells must have the zero value of byte type
			if got.v[i][y] != 0 {
				t.Errorf("field.v[%d][%d] has incorrect value, got: %d, want: %d.", i, y, got.v[i][y], 0)
			}
		}
	}
}

// Test the set value with different coordinates
func TestSetValue(t *testing.T) {
	var x, y int
	var c Coordinate2D
	var e error
	n := 3
	v := byte('X')
	f := newField(n)

	// Coordinate with x < 0
	x, y = -20, 2
	c = *NewCoordinate2D(x, y)
	e = f.setValue(c, v)

	if e == nil {
		t.Error("Set field value with coordinate.x < 0, but no error occured")
	}

	// Coordinate with y < 0
	x, y = 2, -22
	c = *NewCoordinate2D(x, y)
	e = f.setValue(c, v)

	if e == nil {
		t.Error("Set field value with coordinate.y < 0, but no error occured")
	}

	// Coordinate with x > field.n
	x, y = 20, 2
	c = *NewCoordinate2D(x, y)
	e = f.setValue(c, v)

	if e == nil {
		t.Errorf("Set field value with coordinate.x > n(%d), but no error occured", n)
	}

	// Coordinate with x == field.n
	x, y = 3, 2
	c = *NewCoordinate2D(x, y)
	e = f.setValue(c, v)

	if e == nil {
		t.Errorf("Set field value with coordinate.x == n(%d), but no error occured", n)
	}

	// Coordinate with y > field.n
	x, y = 2, 20
	c = *NewCoordinate2D(x, y)
	e = f.setValue(c, v)

	if e == nil {
		t.Errorf("Set field value with coordinate.y > n(%d), but no error occured", n)
	}

	// Coordinate with y == field.n
	x, y = 2, 3
	c = *NewCoordinate2D(x, y)
	e = f.setValue(c, v)

	if e == nil {
		t.Errorf("Set field value with coordinate.y == n(%d), but no error occured", n)
	}

	// Coordinate with 0 < x,y < field.n
	x, y = 2, 2
	c = *NewCoordinate2D(x, y)
	e = f.setValue(c, v)

	if e != nil {
		t.Errorf("Set field value with correct coordinate 0 < x,y < field.n, but an error occured: %v", e)
	}

	if v != f.v[x][y] {
		t.Errorf("f.v[%d][%d] value is incorrect, got: %d, want: %d.", x, y, f.v[x][y], v)
	}
}

// Test if find the cells with value v
func TestFindCellsWithValue(t *testing.T) {
	v, f := byte('X'), newField(3)
	got, vCoordinates := []Coordinate2D{}, []Coordinate2D{
		*NewCoordinate2D(0, 0),
		*NewCoordinate2D(0, 1),
		*NewCoordinate2D(2, 2),
		*NewCoordinate2D(0, 2),
	}

	// Manual set the value v in field with coordinate c
	for _, c := range vCoordinates {
		f.v[c.x][c.y] = v
	}

	got = f.FindCellsWithValue(v)

	if len(got) != len(vCoordinates) {
		t.Errorf("Len of cells found is incorrect, got: %d, want: %d.", len(got), len(vCoordinates))
	}

	for _, c := range vCoordinates {
		f := false

		for j := range got {
			if got[j] == c {
				f = true

				break
			}
		}

		if !f {
			t.Errorf("'%d' v not found in coordinate x:%d y:%d", v, c.x, c.y)
		}
	}
}
