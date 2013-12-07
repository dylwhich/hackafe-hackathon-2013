package board

import (
	"hackathon/ncscreen"
	"hackathon/driver"
)

type Board struct {
	CurrentPosition ncscreen.Coords
	screen ncscreen.Screen
	connection *driver.Connection
}

func NewBoard(position ncscreen.Coords) (*Board, error) {
	conn, err := driver.Connect("/dev/ttyUSB0")
	if err != nil {
		return nil, err
	}

	result := &Board {
		CurrentPosition: position,
		screen: ncscreen.Screen{
			Size: ncscreen.Coords {
				X: 1,
				Y: 1,
			},
			Motors: []ncscreen.Coords{
				ncscreen.Coords{
					X: 0,
					Y: 0,
				},
				ncscreen.Coords{
					X: 1,
					Y: 1,
				},
			},
		},
		connection: conn,
	}

	return result, nil
}

func (b *Board) MoveTo(position ncscreen.Coords){
	var currentDists, targetDists []float64
	currentDists = b.screen.Lengths(b.CurrentPosition)
	
	targetDists = b.screen.Lengths(position)

	motor1move := targetDists[0] - currentDists[0]
	motor2move := targetDists[1] - currentDists[1]

	b.connection.StepDouble(motor1move, motor2move)
}
