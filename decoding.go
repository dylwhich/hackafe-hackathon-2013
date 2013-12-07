package main

import (
	"encoding/json"
	"hackathon/board"
	"io"
)

func DecodeLines(r io.Reader) (lines []*board.Line, err error) {
	lines = make([]*board.Line, 0)
	err = json.NewDecoder(r).Decode(&lines)
	return
}
