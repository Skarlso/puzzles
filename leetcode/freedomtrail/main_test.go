package freedomtrail

import (
	"testing"
)

func TestBasic(t *testing.T) {
	s := findRotateSteps("GODDING", "GD")
	if s != 4 {
		t.Fatal(s)
	}
}

func Test1Basic(t *testing.T) {
	s := findRotateSteps("abcde", "ade")
	if s != 6 {
		t.Fatal(s)
	}
}

func Test2Basic(t *testing.T) {
	// "nyngl"
	// "yyynnnnnnlllggg"
	s := findRotateSteps("nyngl", "yyynnnnnnlllggg")
	if s != 19 {
		t.Fatal("should be 19. was: ", s)
	}
}
