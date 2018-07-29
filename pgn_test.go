package pgn_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/pgn"
)

// var game = `[Event "F/S Return Match"]
// [Site "Belgrade, Serbia JUG"]
// [Date "1992.11.04"]
// [Round "29"]
// [White "Fischer, Robert J."]
// [Black "Spassky, Boris V."]
// [Result "1/2-1/2"]

// 1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}
// 4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8 10. d4 Nbd7
// 11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5
// Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6
// 23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5
// hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5
// 35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6
// Nf2 42. g4 Bd3 43. Re6 1/2-1/2`

var game = `[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]
[Round "29"]
[White "Fischer, Robert J."]
[Black "Spassky, Boris V."]
[Result "1/2-1/2"]

1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}`

func TestParse(t *testing.T) {
	expected := pgn.PGN{
		TagPairs: pgn.TagPairs{
			Event:  "F/S Return Match",
			Site:   "Belgrade, Serbia JUG",
			Date:   "1992.11.04",
			Round:  "29",
			White:  "Fischer, Robert J.",
			Black:  "Spassky, Boris V.",
			Result: "1/2-1/2",
		},
		Movetext: pgn.Movetext{
			pgn.MovetextEntry{White: "e4", Black: "e5"},
			pgn.MovetextEntry{White: "Nf3", Black: "Nc6"},
			pgn.MovetextEntry{White: "Bb5", Black: "a6", Comments: []string{"This opening is called the Ruy Lopez."}},
		},
	}
	got := pgn.Parse(game)
	if got.TagPairs != expected.TagPairs {
		fmt.Println("Got     :", got)
		fmt.Println("Expected: ", expected)
		t.Fatal("failed")
	}
	if !assertMovetextEquality(expected.Movetext, got.Movetext) {
		fmt.Println("Got     :", got.Movetext)
		fmt.Println("Expected: ", expected.Movetext)
		t.Fatal("movetext is unexpected")
	}
}

func TestOuptut(t *testing.T) {
	parsed := pgn.PGN{
		TagPairs: pgn.TagPairs{
			Event:  "F/S Return Match",
			Site:   "Belgrade, Serbia JUG",
			Date:   "1992.11.04",
			Round:  "29",
			White:  "Fischer, Robert J.",
			Black:  "Spassky, Boris V.",
			Result: "1/2-1/2",
		},
		Movetext: pgn.Movetext{
			pgn.MovetextEntry{White: "e4", Black: "e5"},
		},
	}
	var expected = `[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]
[Round "29"]
[White "Fischer, Robert J."]
[Black "Spassky, Boris V."]
[Result "1/2-1/2"]

1. e4 e5`

	got := parsed.String()
	if got != expected {
		fmt.Printf("Got:\n%s\n", got)
		fmt.Printf("Expected:\n%s\n", expected)
		t.Fatal("output failed")
	}

}

func assertMovetextEquality(a, b pgn.Movetext) bool {
	if len(a) != len(b) {
		return false
	}
	for i, valA := range a {
		if valA.White != b[i].White || valA.Black != b[i].Black {
			return false
		}
		if len(valA.Comments) != len(b[i].Comments) {
			return false
		}
		// for j, valB := range valA {
		// 	fmt.Println(valB)
		// 	if valB != b[i][j] {
		// 		return false
		// 	}
		// }
	}
	return true
}
