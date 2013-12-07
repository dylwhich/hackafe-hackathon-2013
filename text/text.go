package text

import (
	"fmt"
	"hackathon/board"
	"hackathon/ncscreen"
	"math"
	"strings"
)

type BoundingBox struct {
	A, B ncscreen.Coords
}

type TextWriter struct {
	font            *Font
	target          *board.Board
	fWidth, fHeight float64
}

func (box *BoundingBox) Size() (width float64, height float64) {
	diff := ncscreen.Subtract(box.B, box.A)
	return math.Abs(diff.X), math.Abs(diff.Y)
}

// Writes the string to the coordinates, regardless of anything except
// out-of-bounds-ness
func (writer *TextWriter) Write(c ncscreen.Coords, text string) {
	text = strings.ToUpper(text)
	println(writer.font)
	for _, char := range text {
		if char == ' ' {
			c.X += 1.2 * writer.fWidth
		} else {
			glyph := writer.font.GetGlyph(char)
			for _, line := range glyph.Lines {
				if line != nil {
					newLine := &board.Line{
						Start: ncscreen.Coords{
							X: c.X + line.Start.X*writer.fWidth,
							Y: c.Y + line.Start.Y*writer.fHeight,
						},
						End: ncscreen.Coords{
							X: c.X + line.End.X*writer.fWidth,
							Y: c.Y + line.End.Y*writer.fHeight,
						},
					}
					fmt.Printf("Line %s %s\n", newLine.Start, newLine.End)

					writer.target.DrawLine(newLine)
				}
			}
			c.X += 1.1 * writer.fWidth
		}
	}
}

func (writer *TextWriter) WriteBoundingBox(bb BoundingBox, text string) {
	w, h := bb.Size()
	writer.WriteCharBox(bb.A, int(h/writer.fHeight), int(w/writer.fWidth), text)
}

func (writer *TextWriter) WriteCharBox(c ncscreen.Coords, rows int,
	cols int, text string) {
	wrapped := strings.Split(Wrap(text, cols), "\n")

	for i, line := range wrapped {
		if i >= rows {
			break
		}

		writer.Write(c, line)
		c.Y += writer.fHeight
	}
}

func NewTextWriter(theFont *Font, charW float64, charH float64, b *board.Board) *TextWriter {
	return &TextWriter{
		font:    theFont,
		target:  b,
		fWidth:  charW,
		fHeight: charH,
	}
}
