package pgn_test

import (
	"fmt"
	"testing"

	pgn "github.com/miketmoore/pgn"
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
		name         string
		in           string
		out          []pgn.Token
		errorMessage string
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
				pgn.Token{Type: pgn.String, Value: "F/S      Return      Match"},
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
			name:         "Movetext - Castle Invalid",
			in:           "1. O-",
			out:          []pgn.Token{},
			errorMessage: pgn.ERR_CASTLE,
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
		{
			name: "Movetext - Castle Queenside - White",
			in:   "18. O-O-O Be7",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "18"},
					pgn.Token{Type: pgn.CastleQueenside, Value: "O-O-O"},
					pgn.Token{Type: pgn.Piece, Value: "B"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "7"},
				},
			),
		},
		{
			name: "Movetext - Castle Queenside - Black",
			in:   "45. Ka6 O-O-O",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "45"},
					pgn.Token{Type: pgn.Piece, Value: "K"},
					pgn.Token{Type: pgn.File, Value: "a"},
					pgn.Token{Type: pgn.Rank, Value: "6"},
					pgn.Token{Type: pgn.CastleQueenside, Value: "O-O-O"},
				},
			),
		},
		{
			name: "Movetext - Castle Queenside - Both",
			in:   "3 O-O-O O-O-O",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "3"},
					pgn.Token{Type: pgn.CastleQueenside, Value: "O-O-O"},
					pgn.Token{Type: pgn.CastleQueenside, Value: "O-O-O"},
				},
			),
		},
		{
			name: "Movetext - Result Draw - Black",
			in:   "43. Re6 1/2-1/2",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "43"},
					pgn.Token{Type: pgn.Piece, Value: "R"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "6"},
					pgn.Token{Type: pgn.Draw, Value: "1/2-1/2"},
				},
			),
		},
		{
			name: "Movetext - Result Draw - White",
			in:   "43. 1/2-1/2 Re6",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "43"},
					pgn.Token{Type: pgn.Draw, Value: "1/2-1/2"},
					pgn.Token{Type: pgn.Piece, Value: "R"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "6"},
				},
			),
		},
		{
			name: "Movetext - Checking Move - White",
			in:   "43. Ke6+ Bf4",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "43"},
					pgn.Token{Type: pgn.Piece, Value: "K"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "6"},
					pgn.Token{Type: pgn.Check, Value: "+"},
					pgn.Token{Type: pgn.Piece, Value: "B"},
					pgn.Token{Type: pgn.File, Value: "f"},
					pgn.Token{Type: pgn.Rank, Value: "4"},
				},
			),
		},
		{
			name: "Movetext - Checking Move - Black",
			in:   "43. Ke6 Bf4+",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "43"},
					pgn.Token{Type: pgn.Piece, Value: "K"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "6"},
					pgn.Token{Type: pgn.Piece, Value: "B"},
					pgn.Token{Type: pgn.File, Value: "f"},
					pgn.Token{Type: pgn.Rank, Value: "4"},
					pgn.Token{Type: pgn.Check, Value: "+"},
				},
			),
		},
		{
			name: "Movetext - Checkmating Move - White",
			in:   "43. Ke6# Bf4",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "43"},
					pgn.Token{Type: pgn.Piece, Value: "K"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "6"},
					pgn.Token{Type: pgn.Checkmate, Value: "#"},
					pgn.Token{Type: pgn.Piece, Value: "B"},
					pgn.Token{Type: pgn.File, Value: "f"},
					pgn.Token{Type: pgn.Rank, Value: "4"},
				},
			),
		},
		{
			name: "Movetext - Checkmating Move - Black",
			in:   "43. Ke6 Bf4#",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "43"},
					pgn.Token{Type: pgn.Piece, Value: "K"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "6"},
					pgn.Token{Type: pgn.Piece, Value: "B"},
					pgn.Token{Type: pgn.File, Value: "f"},
					pgn.Token{Type: pgn.Rank, Value: "4"},
					pgn.Token{Type: pgn.Checkmate, Value: "#"},
				},
			),
		},
		{
			name: "Movetext - Pawn Promotion - White",
			in:   "2. e4= Rf5",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "2"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "4"},
					pgn.Token{Type: pgn.Promotion, Value: "="},
					pgn.Token{Type: pgn.Piece, Value: "R"},
					pgn.Token{Type: pgn.File, Value: "f"},
					pgn.Token{Type: pgn.Rank, Value: "5"},
				},
			),
		},
		{
			name: "Movetext - Pawn Promotion - Black",
			in:   "2. Rf5 e4=",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "2"},
					pgn.Token{Type: pgn.Piece, Value: "R"},
					pgn.Token{Type: pgn.File, Value: "f"},
					pgn.Token{Type: pgn.Rank, Value: "5"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "4"},
					pgn.Token{Type: pgn.Promotion, Value: "="},
				},
			),
		},
		{
			name: "Movetext - Capture - White",
			in:   "2. Rxf5 e4",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "2"},
					pgn.Token{Type: pgn.Piece, Value: "R"},
					pgn.Token{Type: pgn.Capture, Value: "x"},
					pgn.Token{Type: pgn.File, Value: "f"},
					pgn.Token{Type: pgn.Rank, Value: "5"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "4"},
				},
			),
		},
		{
			name: "Movetext - Capture - Black",
			in:   "2. Rf5 xe4",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.MoveNumber, Value: "2"},
					pgn.Token{Type: pgn.Piece, Value: "R"},
					pgn.Token{Type: pgn.File, Value: "f"},
					pgn.Token{Type: pgn.Rank, Value: "5"},
					pgn.Token{Type: pgn.Capture, Value: "x"},
					pgn.Token{Type: pgn.File, Value: "e"},
					pgn.Token{Type: pgn.Rank, Value: "4"},
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
			if test.errorMessage != "" {
				// Expect an error
				if err == nil {
					t.Fatal("Expected an error but did not receive one")
				}
				if err.Error() != test.errorMessage {
					t.Fatal("Unexpected error message found")
				}
			} else {
				// Do not expect error message
				if err != nil {
					fmt.Println("Error: ", err)
					fmt.Println("Tokens: ", tokens)
					t.Fatal("Unexpected error returned")
				}
			}

			if len(tokens) != len(test.out) {
				fmt.Println(tokens)
				t.Fatal("Unexpected total tokens")
			}

			for i, a := range tokens {
				b := test.out[i]
				if a.Type != b.Type || a.Value != b.Value {
					fmt.Printf("Got:\n%v\n", a)
					fmt.Printf("Exp:\n%v\n", b)
					t.Fatal("Unexpected token")
				}
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
