package text

import (
	"encoding/json"
	"errors"
	"hackathon/board"
	"hackathon/ncscreen"
	"os"
)

type Glyph struct {
	Lines     []*board.Line
	Character rune
	Joining   bool
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

	v := make(map[string]interface{}, 0)
	err = json.NewDecoder(file).Decode(&v)
	if err != nil {
		print("Unable to decode json: " + err.Error())
		return nil, err
	}

	font := make(Font, len(v))

	for key, val := range v {
		obj, ok := val.(map[string][]float64)
		if !ok {
			return nil, errors.New("Could not cast to map[string][]float64 at key " + key)
		}

		if len(obj["x"]) != len(obj["y"]) {
			panic("Glyph " + key + " has incompatible lengths")
		}

		lines := make([]*board.Line, 0)
		lineNum := 0

		lastX := -1.0
		lastY := -1.0
		for i := 0; i < len(obj["x"]); i++ {
			x := obj["x"][i]
			y := obj["y"][i]

			if x != -1 && y != -1 {
				if lastX != -1 && lastY != -1 {
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
