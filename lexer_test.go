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
				pgn.Token{Type: pgn.TokenTagName, Value: "Event"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "F/S Return Match"},
			},
		},
		{
			name: "Tag Pair - Lots of Whitespace",
			in:   "  [   Event \"F/S      Return      Match\"    ]    ",
			out: []pgn.Token{
				pgn.Token{Type: pgn.TokenTagName, Value: "Event"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "F/S      Return      Match"},
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
				pgn.Token{Type: pgn.TokenTagName, Value: "Event"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "F/S Return Match"},
				pgn.Token{Type: pgn.TokenTagName, Value: "Site"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "Belgrade, Serbia JUG"},
				pgn.Token{Type: pgn.TokenTagName, Value: "Date"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "1992.11.04"},
				pgn.Token{Type: pgn.TokenTagName, Value: "Round"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "29"},
				pgn.Token{Type: pgn.TokenTagName, Value: "White"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "Fischer, Robert J."},
				pgn.Token{Type: pgn.TokenTagName, Value: "Black"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "Spassky, Boris V."},
				pgn.Token{Type: pgn.TokenTagName, Value: "Result"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "1/2-1/2"},
				pgn.Token{Type: pgn.TokenTagName, Value: "a"},
				pgn.Token{Type: pgn.TokenTagValue, Value: ""},
				pgn.Token{Type: pgn.TokenTagName, Value: "A"},
				pgn.Token{Type: pgn.TokenTagValue, Value: ""},
				pgn.Token{Type: pgn.TokenTagName, Value: "_"},
				pgn.Token{Type: pgn.TokenTagValue, Value: ""},
				pgn.Token{Type: pgn.TokenEOF},
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
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "18"},
					pgn.Token{Type: pgn.TokenCastleKingside, Value: "O-O"},
					pgn.Token{Type: pgn.TokenPiece, Value: "B"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "7"},
				},
			),
		},
		{
			name: "Movetext - Castle Kingside - Black",
			in:   "45. Ka6 O-O",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "45"},
					pgn.Token{Type: pgn.TokenPiece, Value: "K"},
					pgn.Token{Type: pgn.TokenFile, Value: "a"},
					pgn.Token{Type: pgn.TokenRank, Value: "6"},
					pgn.Token{Type: pgn.TokenCastleKingside, Value: "O-O"},
				},
			),
		},
		{
			name: "Movetext - Castle Kingside - Both",
			in:   "3 O-O O-O",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "3"},
					pgn.Token{Type: pgn.TokenCastleKingside, Value: "O-O"},
					pgn.Token{Type: pgn.TokenCastleKingside, Value: "O-O"},
				},
			),
		},
		{
			name: "Movetext - Castle Queenside - White",
			in:   "18. O-O-O Be7",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "18"},
					pgn.Token{Type: pgn.TokenCastleQueenside, Value: "O-O-O"},
					pgn.Token{Type: pgn.TokenPiece, Value: "B"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "7"},
				},
			),
		},
		{
			name: "Movetext - Castle Queenside - Black",
			in:   "45. Ka6 O-O-O",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "45"},
					pgn.Token{Type: pgn.TokenPiece, Value: "K"},
					pgn.Token{Type: pgn.TokenFile, Value: "a"},
					pgn.Token{Type: pgn.TokenRank, Value: "6"},
					pgn.Token{Type: pgn.TokenCastleQueenside, Value: "O-O-O"},
				},
			),
		},
		{
			name: "Movetext - Castle Queenside - Both",
			in:   "3 O-O-O O-O-O",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "3"},
					pgn.Token{Type: pgn.TokenCastleQueenside, Value: "O-O-O"},
					pgn.Token{Type: pgn.TokenCastleQueenside, Value: "O-O-O"},
				},
			),
		},
		{
			name: "Movetext - Result TokenDraw - Black",
			in:   "43. Re6 1/2-1/2",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "43"},
					pgn.Token{Type: pgn.TokenPiece, Value: "R"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "6"},
					pgn.Token{Type: pgn.TokenDraw, Value: "1/2-1/2"},
				},
			),
		},
		{
			name: "Movetext - Result TokenDraw - White",
			in:   "43. 1/2-1/2 Re6",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "43"},
					pgn.Token{Type: pgn.TokenDraw, Value: "1/2-1/2"},
					pgn.Token{Type: pgn.TokenPiece, Value: "R"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "6"},
				},
			),
		},
		{
			name: "Movetext - Checking Move - White",
			in:   "43. Ke6+ Bf4",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "43"},
					pgn.Token{Type: pgn.TokenPiece, Value: "K"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "6"},
					pgn.Token{Type: pgn.TokenCheck, Value: "+"},
					pgn.Token{Type: pgn.TokenPiece, Value: "B"},
					pgn.Token{Type: pgn.TokenFile, Value: "f"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
				},
			),
		},
		{
			name: "Movetext - Checking Move - Black",
			in:   "43. Ke6 Bf4+",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "43"},
					pgn.Token{Type: pgn.TokenPiece, Value: "K"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "6"},
					pgn.Token{Type: pgn.TokenPiece, Value: "B"},
					pgn.Token{Type: pgn.TokenFile, Value: "f"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
					pgn.Token{Type: pgn.TokenCheck, Value: "+"},
				},
			),
		},
		{
			name: "Movetext - Checkmating Move - White",
			in:   "43. Ke6# Bf4",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "43"},
					pgn.Token{Type: pgn.TokenPiece, Value: "K"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "6"},
					pgn.Token{Type: pgn.TokenCheckmate, Value: "#"},
					pgn.Token{Type: pgn.TokenPiece, Value: "B"},
					pgn.Token{Type: pgn.TokenFile, Value: "f"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
				},
			),
		},
		{
			name: "Movetext - Checkmating Move - Black",
			in:   "43. Ke6 Bf4#",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "43"},
					pgn.Token{Type: pgn.TokenPiece, Value: "K"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "6"},
					pgn.Token{Type: pgn.TokenPiece, Value: "B"},
					pgn.Token{Type: pgn.TokenFile, Value: "f"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
					pgn.Token{Type: pgn.TokenCheckmate, Value: "#"},
				},
			),
		},
		{
			name: "Movetext - Pawn Promotion - White",
			in:   "2. e4=B Rf5",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "2"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
					pgn.Token{Type: pgn.TokenPromotionIndicator, Value: "="},
					pgn.Token{Type: pgn.TokenPromotionPiece, Value: "B"},
					pgn.Token{Type: pgn.TokenPiece, Value: "R"},
					pgn.Token{Type: pgn.TokenFile, Value: "f"},
					pgn.Token{Type: pgn.TokenRank, Value: "5"},
				},
			),
		},
		{
			name: "Movetext - Pawn Promotion - Black",
			in:   "2. Rf5 e4=Q",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "2"},
					pgn.Token{Type: pgn.TokenPiece, Value: "R"},
					pgn.Token{Type: pgn.TokenFile, Value: "f"},
					pgn.Token{Type: pgn.TokenRank, Value: "5"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
					pgn.Token{Type: pgn.TokenPromotionIndicator, Value: "="},
					pgn.Token{Type: pgn.TokenPromotionPiece, Value: "Q"},
				},
			),
		},
		{
			name: "Movetext - TokenCapture - White",
			in:   "2. Rxf5 e4",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "2"},
					pgn.Token{Type: pgn.TokenPiece, Value: "R"},
					pgn.Token{Type: pgn.TokenCapture, Value: "x"},
					pgn.Token{Type: pgn.TokenFile, Value: "f"},
					pgn.Token{Type: pgn.TokenRank, Value: "5"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
				},
			),
		},
		{
			name: "Movetext - TokenCapture - Black",
			in:   "2. Rf5 xe4",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "2"},
					pgn.Token{Type: pgn.TokenPiece, Value: "R"},
					pgn.Token{Type: pgn.TokenFile, Value: "f"},
					pgn.Token{Type: pgn.TokenRank, Value: "5"},
					pgn.Token{Type: pgn.TokenCapture, Value: "x"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
				},
			),
		},
		{
			name:         "Comment without closing brace",
			in:           "1. a4 e5 { aslks  klasdf i23lk43nncklj3#$1412kfdlsjf ",
			out:          []pgn.Token{},
			errorMessage: pgn.ERR_COMMENT_NOT_CLOSED,
		},
		{
			name: "Disambiguation - Originating TokenFile of Moving TokenPiece - White",
			in:   "10. Nbd7 d4",
			out: []pgn.Token{
				pgn.Token{Type: pgn.TokenMoveNumber, Value: "10"},
				pgn.Token{Type: pgn.TokenPiece, Value: "N"},
				pgn.Token{Type: pgn.TokenFile, Value: "b"},
				pgn.Token{Type: pgn.TokenFile, Value: "d"},
				pgn.Token{Type: pgn.TokenRank, Value: "7"},
				pgn.Token{Type: pgn.TokenFile, Value: "d"},
				pgn.Token{Type: pgn.TokenRank, Value: "4"},
			},
		},
		{
			name: "Disambiguation - Originating TokenFile of Moving TokenPiece - Black",
			in:   "10. d4 Nbd7",
			out: []pgn.Token{
				pgn.Token{Type: pgn.TokenMoveNumber, Value: "10"},
				pgn.Token{Type: pgn.TokenFile, Value: "d"},
				pgn.Token{Type: pgn.TokenRank, Value: "4"},
				pgn.Token{Type: pgn.TokenPiece, Value: "N"},
				pgn.Token{Type: pgn.TokenFile, Value: "b"},
				pgn.Token{Type: pgn.TokenFile, Value: "d"},
				pgn.Token{Type: pgn.TokenRank, Value: "7"},
			},
		},
		{
			name: "newline separation between tag pairs and movetext",
			in: `[Event "Blah Blah"]

1. e4 e5`,
			out: []pgn.Token{
				pgn.Token{Type: pgn.TokenTagName, Value: "Event"},
				pgn.Token{Type: pgn.TokenTagValue, Value: "Blah Blah"},
				pgn.Token{Type: pgn.TokenMoveNumber, Value: "1"},
				pgn.Token{Type: pgn.TokenFile, Value: "e"},
				pgn.Token{Type: pgn.TokenRank, Value: "4"},
				pgn.Token{Type: pgn.TokenFile, Value: "e"},
				pgn.Token{Type: pgn.TokenRank, Value: "5"},
			},
		},
		{
			name: "Whole game",
			in: "[Event \"F/S Return Match\"]\n" +
				"[Site \"Belgrade, Serbia JUG\"]\n" +
				"[Date \"1992.11.04\"]\n[Round \"29\"]\n" +
				"[White \"Fischer, Robert J.\"]\n" +
				"[Black \"Spassky, Boris V.\"]\n" +
				"[Result \"1/2-1/2\"]\n" +
				"[a \"\"]\n" +
				"[A \"\"]\n" +
				"\n" +
				"1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}" +
				"4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8 10. d4 Nbd7" +
				"11. c4 c6 12. cxb5 axb5",
			//13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5",
			// "Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6" +
			// "23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5" +
			// "hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5" +
			// "35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6" +
			// "Nf2 42. g4 Bd3 43. Re6 1/2-1/2",
			out: buildTokens(
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenTagName, Value: "Event"},
					pgn.Token{Type: pgn.TokenTagValue, Value: "F/S Return Match"},
					pgn.Token{Type: pgn.TokenTagName, Value: "Site"},
					pgn.Token{Type: pgn.TokenTagValue, Value: "Belgrade, Serbia JUG"},
					pgn.Token{Type: pgn.TokenTagName, Value: "Date"},
					pgn.Token{Type: pgn.TokenTagValue, Value: "1992.11.04"},
					pgn.Token{Type: pgn.TokenTagName, Value: "Round"},
					pgn.Token{Type: pgn.TokenTagValue, Value: "29"},
					pgn.Token{Type: pgn.TokenTagName, Value: "White"},
					pgn.Token{Type: pgn.TokenTagValue, Value: "Fischer, Robert J."},
					pgn.Token{Type: pgn.TokenTagName, Value: "Black"},
					pgn.Token{Type: pgn.TokenTagValue, Value: "Spassky, Boris V."},
					pgn.Token{Type: pgn.TokenTagName, Value: "Result"},
					pgn.Token{Type: pgn.TokenTagValue, Value: "1/2-1/2"},
					pgn.Token{Type: pgn.TokenTagName, Value: "a"},
					pgn.Token{Type: pgn.TokenTagValue, Value: ""},
					pgn.Token{Type: pgn.TokenTagName, Value: "A"},
					pgn.Token{Type: pgn.TokenTagValue, Value: ""},
					pgn.Token{Type: pgn.TokenTagName, Value: "_"},
					pgn.Token{Type: pgn.TokenTagValue, Value: ""},
				},
				newMove("1", "", "e4", "", "e5"),
				newMove("2", "N", "f3", "N", "c6"),
				newMove("3", "B", "b5", "", "a6"),
				newMove("4", "B", "a4", "N", "f6"),
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "5"},
					pgn.Token{Type: pgn.TokenCastleKingside, Value: "O-O"},
					pgn.Token{Type: pgn.TokenPiece, Value: "B"},
					pgn.Token{Type: pgn.TokenFile, Value: "e"},
					pgn.Token{Type: pgn.TokenRank, Value: "7"},
				},
				newMove("6", "R", "e1", "", "b5"),
				newMove("7", "B", "b3", "", "d6"),
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "8"},
					pgn.Token{Type: pgn.TokenFile, Value: "c"},
					pgn.Token{Type: pgn.TokenRank, Value: "3"},
					pgn.Token{Type: pgn.TokenCastleKingside, Value: "O-O"},
				},
				newMove("9", "", "h3", "N", "b8"),
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "10"},
					pgn.Token{Type: pgn.TokenFile, Value: "d"},
					pgn.Token{Type: pgn.TokenRank, Value: "4"},
					pgn.Token{Type: pgn.TokenPiece, Value: "N"},
					pgn.Token{Type: pgn.TokenFile, Value: "b"},
					pgn.Token{Type: pgn.TokenFile, Value: "d"},
					pgn.Token{Type: pgn.TokenRank, Value: "7"},
				},
				newMove("11", "", "c4", "", "c6"),
				[]pgn.Token{
					pgn.Token{Type: pgn.TokenMoveNumber, Value: "12"},
					pgn.Token{Type: pgn.TokenFile, Value: "c"},
					pgn.Token{Type: pgn.TokenCapture, Value: "x"},
					pgn.Token{Type: pgn.TokenFile, Value: "b"},
					pgn.Token{Type: pgn.TokenRank, Value: "5"},
					pgn.Token{Type: pgn.TokenFile, Value: "a"},
					pgn.Token{Type: pgn.TokenCapture, Value: "x"},
					pgn.Token{Type: pgn.TokenFile, Value: "b"},
					pgn.Token{Type: pgn.TokenRank, Value: "5"},
				},
			),
		},
	}

	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			scanner := pgn.NewScanner(test.in)
			lexer := pgn.NewLexer(scanner)

			err, got := lexer.Tokenize()
			if test.errorMessage != "" {
				// Expect an error
				if err == nil {
					t.Fatal("Expected an error but did not receive one")
				}
				if err.Error() != test.errorMessage {
					fmt.Println(err)
					t.Fatal("Unexpected error message found")
				}
			} else {
				// Do not expect error message
				if err != nil {
					fmt.Println("Error: ", err)
					fmt.Println("Tokens: ", got)
					t.Fatal("Unexpected error returned")
				}
			}

			if len(got) != len(test.out) {
				fmt.Println(got)
				t.Fatal("Unexpected total tokens")
			}

			for i, expectedToken := range test.out {
				gotToken := got[i]
				if expectedToken.Type != gotToken.Type || expectedToken.Value != gotToken.Value {
					fmt.Printf("Got:\n%v\n", gotToken)
					fmt.Printf("Exp:\n%v\n", expectedToken)
					t.Fatal("Unexpected token")
				}
			}
		})
	}
}

// func newMove("3", "B", "b5", "", "a6"),
func newMove(n, pA, moveA, pB, moveB string) []pgn.Token {
	tokens := []pgn.Token{
		pgn.Token{Type: pgn.TokenMoveNumber, Value: n},
	}

	if pA != "" {
		tokens = append(tokens, pgn.Token{Type: pgn.TokenPiece, Value: pA})
	}

	tokens = append(tokens, pgn.Token{Type: pgn.TokenFile, Value: string(moveA[0])})
	tokens = append(tokens, pgn.Token{Type: pgn.TokenRank, Value: string(moveA[1])})

	if pB != "" {
		tokens = append(tokens, pgn.Token{Type: pgn.TokenPiece, Value: pB})
	}

	tokens = append(tokens, pgn.Token{Type: pgn.TokenFile, Value: string(moveB[0])})
	tokens = append(tokens, pgn.Token{Type: pgn.TokenRank, Value: string(moveB[1])})

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
