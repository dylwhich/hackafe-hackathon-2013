package ncscreen

import (
	"fmt"
	"math"
)

type Screen struct {
	Size   Coords
	Motors []Coords
}

type Coords struct {
	X, Y float64
}

// Lengths determines the lengths of slack required of each motor string from
// their mount points in order to center them on that point.
func (s *Screen) Lengths(c Coords) []float64 {
	lengths := make([]float64, len(s.Motors))
	for i, motor := range s.Motors {
		lengths[i] = Distance(motor, c)
	}
	return lengths
}

// InBounds checks whether a given coordinate is within the reach of
// the listed motors and the screen itself. The only upper bound on the
// Y coordinate (that is, toward the floor) is considered to be the
// screen boundary.
func (s *Screen) InBounds(c Coords) bool {
	// If the given coordinates are off the screen in any direction,
	// then it is out of bounds, of course.
	if c.X < 0 || c.X > s.Size.X {
		return false
	}

	if c.Y < 0 || c.Y > s.Size.Y {
		return false
	}

	// As long as there is at least one motor above the coordinate
	// position, then that does not cause it to be out of bounds. The
	// coordinates must also be between two motors in the X
	// coordinate.
	hasY := false
	hasLowX := false
	hasHighX := false
	for _, motor := range s.Motors {
		if c.Y >= motor.Y {
			hasY = true
		}

		if c.X >= motor.X {
			hasLowX = true
		}
		if c.X <= motor.X {
			hasHighX = true
		}
	}

	return hasY && hasLowX && hasHighX
}

// String converts the coordinates to a standard string format.
func (c Coords) String() string {
	return fmt.Sprintf("(%f, %f)", c.X, c.Y)
}

func Distance(c1, c2 Coords) float64 {
	return math.Sqrt(math.Pow(c1.X-c2.X, 2) + math.Pow(c1.Y-c2.Y, 2))
}

func Subtract(c1, c2 Coords) (result Coords) {
	result.X = c1.X - c2.X
	result.Y = c1.Y - c2.Y
	return
}

func Add(c1, c2 Coords) (result Coords) {
	result.X = c1.X + c2.X
	result.Y = c1.Y + c2.Y
	return
}
