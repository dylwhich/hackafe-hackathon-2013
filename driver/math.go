package driver

import (
	"math"
)

const (
	// RadiansPerStep is dependent on the motors used on the board.
	RadiansPerStep = math.Pi / 100

	// MotorRadius is dependent on the motors used on the board.
	MotorRadius = 0.004 // 6 cm

	// MotorSteps is the number of steps per turn.
	MotorSteps = 200
)

// TurnsIn converts an arbitrary distance to a number of complete
// turns required.
func TurnsIn(meters float64) float64 {
	return meters / (MotorRadius * 2 * math.Pi)
}

// Steps converts a number of turns to a number of steps.
func Steps(turns float64) int {
	return int(MotorSteps * turns)
}
