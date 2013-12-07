package main

import (
	"flag"
	"fmt"
	"hackathon/board"
	"hackathon/ncscreen"
	"os"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", err)
		return
	}
	defer file.Close()

	lines, err := DecodeLines(file)
	if err != nil {
		fmt.Printf("Could not decode: %s\n", err)
		return
	}

	// Try to open a connection to the board.
	b, err := board.NewBoard(ncscreen.Coords{0.5, 0.5}, false)
	if err != nil {
		fmt.Printf("Could not connect to board: %s\n", err)
		//return
	}

	for _, line := range lines {
		fmt.Printf("Drawing line from %s to %s\n", line.Start, line.End)
		if b != nil {
			b.DrawLine(line)
		}
	}
}
