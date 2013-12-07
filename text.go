package text

import (
	"encoding/json"
	"hackathon/board"
	"hackathon/ncscreen"
	"math"
)

type BoundingBox struct {
	A, B Coords
}

type TextWriter struct {
	font            *Font
	target          Board
	fWidth, fHeight float64
}

func (box *BoundingBox) Size() (float64, float64) {
	diff := board.Subtract(box.B, box.A)
	return math.Abs(diff.X), math.Abs(diff.Y)
}

// Writes the string to the coordinates, regardless of anything except out-of-bounds-ness
func (writer *TextWriter) Write(c Coords, text string) {
	for char := range text {
		glyph := font.GetGlyph(char)
		for line := range glyph.Lines {
			line.Start.X += c.X
			line.End.X += c.X
			line.Start.Y += c.Y
			line.End.Y += c.Y

			board.DrawLine(line)
		}
		c.X += fWidth
	}
}

func (writer *TextWriter) WriteBoundingBox(bb BoundingBox, text string) {
	w, h := bb.Size()
	WriteCharBox(bb.A, int(h/fHeight), int(w/fWidth), text)
}

func (writer *TextWriter) WriteCharBox(c Coords, rows int, cols int, text string) {
	wrapped := string.Split(Wrap(text, cols), "\n")

	for i, line := range wrapped {
		if i >= rows {
			break
		}

		writer.Write(c, line)
		c.Y += fHeight
	}
}

func NewTextWriter(font *Font, charW float64, charH float64, board Board) *TextWriter {
	return TextWriter{
		font:    font,
		target:  board,
		fWidth:  charW,
		fHeight: charH,
	}
}
