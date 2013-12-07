package board

import (
	"hackathon/ncscreen"
)

type Line struct {
	Start, End ncscreen.Coords
}

func (b *Board) DrawLine(l *Line) {
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
