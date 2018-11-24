package pgn_test

import (
	"testing"

	pgn "github.com/miketmoore/pgn-3"
)

func TestScannerPeek(t *testing.T) {
	s := pgn.NewScanner("a")
	ok, r := s.Peek()
	if ok == false || r != 'a' {
		t.Fatal("Peek failed")
	}
}

func TestScannerNext(t *testing.T) {
	s := pgn.NewScanner("ab")
	ok, runeA := s.Next()
	if ok == false || runeA != 'a' {
		t.Fatal("Next failed")
	}
	ok, runeB := s.Next()
	if ok == false || runeB != 'b' {
		t.Fatal("Next failed")
	}
	ok, _ = s.Next()
	if ok == true {
		t.Fatal("Next failed")
	}
}
