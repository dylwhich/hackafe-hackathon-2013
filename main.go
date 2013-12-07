package main

import (
	"flag"
	"fmt"
	"hackathon/board"
	"hackathon/ncscreen"
	"hackathon/text"
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

	/*lines, err := DecodeLines(file)
	if err != nil {
		fmt.Printf("Could not decode: %s\n", err)
		return
	}*/

	// Try to open a connection to the board.
	b, err := board.NewBoard(ncscreen.Coords{0.43, 0.2286}, false)
	if err != nil {
		fmt.Printf("Could not connect to board: %s\n", err)
		//return
	}

	/*for _, line := range lines {
		fmt.Printf("Drawing line from %s to %s\n", line.Start, line.End)
		if b != nil {
			b.DrawLine(line)
		}
	}*/

	theFont, err := text.LoadFont("text/font.json")

	if err != nil {
		fmt.Printf("Unable to load font: %s", err.Error())
	} else {
		writer := text.NewTextWriter(theFont, 1/39.3701, 1.5*1/39.3701, b)

		str := "HELLO, WORLD!\n "
		strCoord := ncscreen.Coords{0.43, 0.2286}
		fmt.Printf("Writing \"%s\" at %s\n", str, strCoord)

		writer.Write(strCoord, str)

	}
}
