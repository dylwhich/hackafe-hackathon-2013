package text

import (
	"testing"
)

func TestLoadFont(t *testing.T) {
	result, err := LoadFont("font.json")
	if err != nil {
		t.Fail()
	}

	if result != nil {
		print("Success!")
	}
}
