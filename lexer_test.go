package pgn_test

import (
	"fmt"
	"testing"

	pgn "github.com/miketmoore/pgn-3"
)

var gameA string = `[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]
[Round "29"]
[White "Fischer, Robert J."]
[Black "Spassky, Boris V."]
[Result "1/2-1/2"]

1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}
4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8 10. d4 Nbd7
11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5
Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6
23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5
hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5
35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6
Nf2 42. g4 Bd3 43. Re6 1/2-1/2`

func TestTokenize(t *testing.T) {
	data := []struct {
		name string
		in   string
		out  []pgn.Token
	}{
		{
			name: "Tag Pair",
			in:   "[Event \"F/S Return Match\"]",
			out: []pgn.Token{
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Event"},
				pgn.Token{Type: pgn.String, Value: "F/S Return Match"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
			},
		},
		{
			name: "Tag Pair - Lots of Whitespace",
			in:   "  [   Event \"F/S      Return      Match\"    ]    ",
			out: []pgn.Token{
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Event"},
				pgn.Token{Type: pgn.String, Value: "F/S Return Match"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
			},
		},
		{
			name: "Multiple Tag Pairs",
			in: "[Event \"F/S Return Match\"]\n" +
				"[Site \"Belgrade, Serbia JUG\"]\n" +
				"[Date \"1992.11.04\"]\n[Round \"29\"]\n" +
				"[White \"Fischer, Robert J.\"]\n" +
				"[Black \"Spassky, Boris V.\"]\n" +
				"[Result \"1/2-1/2\"]\n" +
				"[a \"\"]\n" +
				"[A \"\"]\n" +
				"[_ \"\"]\n",
			out: []pgn.Token{
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Event"},
				pgn.Token{Type: pgn.String, Value: "F/S Return Match"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Site"},
				pgn.Token{Type: pgn.String, Value: "Belgrade, Serbia JUG"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Date"},
				pgn.Token{Type: pgn.String, Value: "1992.11.04"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Round"},
				pgn.Token{Type: pgn.String, Value: "29"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "White"},
				pgn.Token{Type: pgn.String, Value: "Fischer, Robert J."},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Black"},
				pgn.Token{Type: pgn.String, Value: "Spassky, Boris V."},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "Result"},
				pgn.Token{Type: pgn.String, Value: "1/2-1/2"},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "a"},
				pgn.Token{Type: pgn.String, Value: ""},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "A"},
				pgn.Token{Type: pgn.String, Value: ""},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
				pgn.Token{Type: pgn.TagPairOpen, Value: "["},
				pgn.Token{Type: pgn.TagName, Value: "_"},
				pgn.Token{Type: pgn.String, Value: ""},
				pgn.Token{Type: pgn.TagPairClose, Value: "]"},
			},
		},
		{
			name: "Movetext - One",
			in:   "1. e4 e5",
			out: buildTokens(
				newMove("1", "", "e4", "", "e5"),
			),
		},
		{
			name: "Movetext - Two",
			in:   "1. e4 e5 2. Nf3 Nc6",
			out: buildTokens(
				newMove("1", "", "e4", "", "e5"),
				newMove("2", "N", "f3", "N", "c6"),
			),
		},
		{
			name: "Movetext - Castle Kingside - White",
			in:   "18. O-O Be7",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "18"},
					pgn.Token{Type: pgn.CastleKingside, Value: "O-O"},
					pgn.Token{Type: pgn.Piece, Value: "B"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "7"},
				},
			),
		},
		{
			name: "Movetext - Castle Kingside - Black",
			in:   "45. Ka6 O-O",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "45"},
					pgn.Token{Type: pgn.Piece, Value: "K"},
					pgn.Token{Type: pgn.File, Value: "a"},
					pgn.Token{Type: pgn.Rank, Value: "6"},
					pgn.Token{Type: pgn.CastleKingside, Value: "O-O"},
				},
			),
		},
		{
			name: "Movetext - Castle Kingside - Both",
			in:   "3 O-O O-O",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "3"},
					pgn.Token{Type: pgn.CastleKingside, Value: "O-O"},
					pgn.Token{Type: pgn.CastleKingside, Value: "O-O"},
				},
			),
		},
		// {
		// 	name: "Whole game of movetext",
		// 	in: "1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}" +
		// 		"4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8",
		// 	out: buildTokens(
		// 		newMove("1", "", "e4", "", "e5"),
		// 		newMove("2", "N", "f3", "N", "c6"),
		// 		newMove("3", "B", "b5", "", "a6"),
		// 		newMove("4", "B", "a4", "N", "f6"),
		// 		[]pgn.Token{
		// 			pgn.Token{Type: pgn.MoveNumber, Value: "5"},
		// 			pgn.Token{Type: pgn.CastleKingside, Value: "O-O"},
		// 			pgn.Token{Type: pgn.Piece, Value: "B"},
		// 			pgn.Token{Type: pgn.File, Value: "e"},
		// 			pgn.Token{Type: pgn.Rank, Value: "7"},
		// 		},
		// 		newMove("6", "R", "e1", "", "b5"),
		// 		newMove("7", "B", "b3", "", "d6"),
		// 		[]pgn.Token{
		// 			pgn.Token{Type: pgn.MoveNumber, Value: "8"},
		// 			pgn.Token{Type: pgn.File, Value: "c"},
		// 			pgn.Token{Type: pgn.Rank, Value: "3"},
		// 			pgn.Token{Type: pgn.CastleKingside, Value: "O-O"},
		// 		},
		// 		newMove("9", "", "h3", "N", "b8"),
		// 	),
		// },
	}

	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			scanner := pgn.NewScanner(test.in)
			lexer := pgn.NewLexer(scanner)

			tokens := []pgn.Token{}
			err, tokens := lexer.Tokenize(tokens)
			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println("Tokens: ", tokens)
				t.Fatal("Unexpected error returned")
			}
			if len(tokens) != len(test.out) {
				fmt.Println(tokens)
				t.Fatal("Unexpected total tokens")
			}
		})
	}
}

// func newMove("3", "B", "b5", "", "a6"),
func newMove(n, pA, moveA, pB, moveB string) []pgn.Token {
	tokens := []pgn.Token{
		pgn.Token{Type: pgn.MoveNumber, Value: n},
	}

	if pA != "" {
		tokens = append(tokens, pgn.Token{Type: pgn.Piece, Value: pA})
	}

	tokens = append(tokens, pgn.Token{Type: pgn.File, Value: string(moveA[0])})
	tokens = append(tokens, pgn.Token{Type: pgn.Rank, Value: string(moveA[1])})

	if pB != "" {
		tokens = append(tokens, pgn.Token{Type: pgn.Piece, Value: pB})
	}

	tokens = append(tokens, pgn.Token{Type: pgn.File, Value: string(moveB[0])})
	tokens = append(tokens, pgn.Token{Type: pgn.Rank, Value: string(moveB[1])})

	return tokens
}

func buildTokens(moveTokens ...[]pgn.Token) []pgn.Token {
	tokens := []pgn.Token{}
	for _, mt := range moveTokens {
		for _, t := range mt {
			tokens = append(tokens, t)
		}
	}
	return tokens
}
