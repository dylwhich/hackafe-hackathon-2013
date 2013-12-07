package text

import (
	"testing"
)

func TestLoadFont(t *testing.T) {
	result := LoadFont("font.json")
	if result != nil {
		print("Success?")
	}
}
