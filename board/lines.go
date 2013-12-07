package board

import (
	"hackathon/ncscreen"
	"math"
)

type Line struct {
	Start, End ncscreen.Coords
}

func (l *Line) Split(pieces int) []Line {
	result := make([]Line, pieces)

	dy := l.End.Y - l.Start.Y
	dx := l.End.X - l.Start.X

	length := ncscreen.Distance(l.Start, l.End) / pieces

	theta := math.Atan2(dy, dx)

	for i := 0; i < pieces; i++ {
		result[i] = Line{
			Start: ncscreen.Coords{
				X: float64(i) * length * math.Cos(theta),
				Y: float64(i) * length * math.Sin(theta),
			},
			End: ncscreen.Coords{
				X: float64(i+1) * length * math.Cos(theta),
				Y: float64(i+1) * length * math.Sin(theta),
			},
		}
	}
	return result
}

func (b *Board) DrawLine(l *Line) {
	for _, segment := range l.Split(50) {
		// First, make sure that there we don't draw where we don't mean
		// to.
		b.SetPenDown(false)

		// Find the nearest end of the line and move the marker to there.
		near, far := segment.FindNearerTo(b.CurrentPosition)
		b.MoveTo(near)

		// Draw the line.
		b.SetPenDown(true)
		b.MoveTo(far)
		b.SetPenDown(false)
	}
}

func (l *Line) FindNearerTo(c ncscreen.Coords) (near, far ncscreen.Coords) {
	distStart := ncscreen.Distance(c, l.Start)
	distEnd := ncscreen.Distance(c, l.End)

	if distStart <= distEnd {
		return l.Start, l.End
	} else {
		return l.End, l.Start
	}
}
