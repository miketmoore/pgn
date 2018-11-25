package pgn_test

import (
	"testing"

	pgn "github.com/miketmoore/pgn-3"
)

func TestScannerPeek(t *testing.T) {
	s := pgn.NewScanner("a")
	r := s.Peek()
	if r != 'a' {
		t.Fatal("Peek failed")
	}
}

func TestScannerNext(t *testing.T) {
	s := pgn.NewScanner("ab")
	runeA := s.Next()
	if runeA != 'a' {
		t.Fatal("Next failed")
	}
	runeB := s.Next()
	if runeB != 'b' {
		t.Fatal("Next failed")
	}
	runeC := s.Next()
	if runeC != rune(0) {
		t.Fatal("Next failed")
	}
}
