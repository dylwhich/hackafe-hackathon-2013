package board

import (
	"fmt"
	"hackathon/driver"
	"hackathon/ncscreen"
)

type Board struct {
	CurrentPosition ncscreen.Coords
	PenDown         bool
	screen          ncscreen.Screen
	connection      *driver.Connection
}

func NewBoard(position ncscreen.Coords, penDown bool) (*Board, error) {
	conn, err := driver.Connect("/dev/ttyUSB0")
	if err != nil {
		return nil, err
	}

	result := &Board{
		CurrentPosition: position,
		PenDown:         penDown,
		screen: ncscreen.Screen{
			Size: ncscreen.Coords{
				X: .9625,
				Y: .508,
			},
			Motors: []ncscreen.Coords{
				ncscreen.Coords{
					X: 0,
					Y: -.2286,
				},
				ncscreen.Coords{
					X: .9625,
					Y: -.2286,
				},
			},
		},
		connection: conn,
	}

	return result, nil
}

func (b *Board) SetPenDown(penDown bool) {
	if penDown != b.PenDown {
		b.PenDown = penDown
		if penDown {
			b.connection.PenDown()
		} else {
			b.connection.PenUp()
		}
	}
}

func (b *Board) MoveTo(position ncscreen.Coords) {
	var currentDists, targetDists []float64
	currentDists = b.screen.Lengths(b.CurrentPosition)

	targetDists = b.screen.Lengths(position)

	motor1move := targetDists[0] - currentDists[0]
	motor2move := targetDists[1] - currentDists[1]

	b.connection.MoveRelativeDouble(motor1move, motor2move)

	fmt.Printf("I'm at %s\n", b.CurrentPosition)

	b.CurrentPosition = position

	fmt.Printf("I'm moving to %s\n", b.CurrentPosition)
}
