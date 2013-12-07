package ncscreen

import (
	"math"
	"testing"
)

func TestLengths(t *testing.T) {
	s := Screen{
		Size: Coords{10, 10},
		Motors: []Coords{
			{0, 0},
			{10, 0},
		},
	}

	lengths := s.Lengths(Coords{0, 0})
	if lengths[0] != float64(0) || lengths[1] != float64(10) {
		t.Errorf("Lengths %s != [0, 10]\n", lengths)
		t.Fail()
	}

	lengths = s.Lengths(Coords{5, 5})
	sqrt50 := math.Sqrt(float64(50))
	if lengths[0] != sqrt50 || lengths[1] != sqrt50 {
		t.Errorf("Lengths %s != [sqrt(50), sqrt(50)]\n", lengths)
		t.Fail()
	}
}

func TestInBounds(t *testing.T) {
	s := Screen{
		Size: Coords{10, 10},
		Motors: []Coords{
			{0, 0},
			{9, 0},
		},
	}

	// Testing in bounds.

	if !s.InBounds(Coords{0, 0}) {
		t.Errorf("(0, 0) should be in bounds, but is not")
		t.Fail()
	}

	if !s.InBounds(Coords{9, 0}) {
		t.Errorf("(9, 0) should be in bounds, but is not")
		t.Fail()
	}

	if !s.InBounds(Coords{5, 10}) {
		t.Errorf("(5, 10) should be in bounds, but is not")
		t.Fail()
	}

	// Testing out of bounds.

	if s.InBounds(Coords{5, -1}) {
		t.Errorf("(5, -1) should be out of bounds, but is not")
		t.Fail()
	}

	if s.InBounds(Coords{5, 11}) {
		t.Errorf("(5, 11) should be out of bounds, but is not")
		t.Fail()
	}

	if s.InBounds(Coords{10, 5}) {
		t.Errorf("(10, 5) should be out of bounds, but is not")
		t.Fail()
	}
}
