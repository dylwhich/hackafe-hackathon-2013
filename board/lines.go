package board

import (
	"hackathon/ncscreen"
)

type Line struct {
	Start, End ncscreen.Coords
}

func (l *Line) Split(pieces int) []Line {
	result := make([]Line, pieces)

	dy := l.End.Y - l.Start.Y
	dx := l.End.X - l.Start.X

	length := ncscreen.Distance(l.Start, l.End)

	normVector := ncscreen.Coords{dx / length, dy / length}

	for i := 0; i < pieces; i++ {
		result[i] = Line{
			Start: ncscreen.Add(l.Start, ncscreen.Coords{normVector.X * float64(i), normVector.Y * float64(i)}),
			End:   ncscreen.Add(l.Start, ncscreen.Coords{normVector.X * float64(i+1), normVector.Y * float64(i+1)}),
		}
	}

	return result
}

func (b *Board) DrawLine(l *Line) {
	//for _, segment := range l.Split(50) {
	// First, make sure that there we don't draw where we don't mean
	// to.
	b.SetPenDown(false)

	// Find the nearest end of the line and move the marker to there.
	near, far := l.FindNearerTo(b.CurrentPosition)
	b.MoveTo(near)

	// Draw the line.
	b.SetPenDown(true)
	b.MoveTo(far)
	b.SetPenDown(false)
	//}
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
