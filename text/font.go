package text

import (
	"encoding/json"
	"hackathon/board"
	"hackathon/ncscreen"
	"os"
)

type Glyph struct {
	Lines     []*board.Line
	Character rune
	Joining   bool
}

type TmpChar map[string]struct {
	X []float64 `json:"x"`
	Y []float64 `json:"y"`
}

type Font map[rune]Glyph

func (f Font) GetGlyph(c rune) Glyph {
	result, ok := f[c]
	if !ok {
		result = f[rune('_')]
	}

	return result
}

func LoadFont(fileName string) (result *Font, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		print("Error opening " + fileName + ": " + err.Error())
		return nil, err
	}
	defer file.Close()

	v := make(TmpChar, 0)
	err = json.NewDecoder(file).Decode(&v)
	if err != nil {
		print("Unable to decode json: " + err.Error())
		return nil, err
	}

	font := make(Font, len(v))

	for key, val := range v {
		lines := make([]*board.Line, len(val.X))
		lineNum := 0

		lastX := -1.0
		lastY := -1.0
		for i := 0; i < len(val.X); i++ {
			x := val.X[i]
			y := val.Y[i]

			if x != -1.0 && y != -1.0 {
				if lastX != -1.0 && lastY != -1.0 {
					lines[lineNum] = &board.Line{
						Start: ncscreen.Coords{
							X: lastX,
							Y: lastY,
						},
						End: ncscreen.Coords{
							X: x,
							Y: y,
						},
					}

					lineNum++
				}

				lastX = x
				lastY = y
			}

		}

		font[rune(key[0])] = Glyph{
			Lines:     lines,
			Character: rune(key[0]),
			Joining:   false,
		}
	}

	return &font, nil
}
