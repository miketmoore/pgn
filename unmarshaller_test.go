package pgn_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/pgn"
)

func TestUnmarshal(t *testing.T) {
	data := []struct {
		name         string
		in           string
		out          pgn.PGN
		errorMessage string
	}{
		{
			name: "Empty game",
			in:   "",
			out: pgn.PGN{
				Games: []pgn.Game{
					pgn.Game{},
				},
			},
		},
		{
			name: "Tag pair",
			in: "[Event \"F/S Return Match\"]\n" +
				"[Site \"Belgrade, Serbia JUG\"]\n" +
				"[Date \"1992.11.04\"]\n[Round \"29\"]\n" +
				"[White \"Fischer, Robert J.\"]\n" +
				"[Black \"Spassky, Boris V.\"]\n" +
				"[Result \"1/2-1/2\"]\n" +
				"[a \"\"]\n" +
				"[A \"\"]\n" +
				"[_ \"\"]\n",
			out: pgn.PGN{
				Games: []pgn.Game{
					pgn.Game{
						TagPairs: []pgn.TagPair{
							pgn.TagPair{Name: "Event", Value: "F/S Return Match"},
							pgn.TagPair{Name: "Site", Value: "Belgrad"},
							pgn.TagPair{Name: "Date", Value: "1992.11.04"},
							pgn.TagPair{Name: "Round", Value: "29"},
							pgn.TagPair{Name: "White", Value: "Fischer, Robert J."},
							pgn.TagPair{Name: "Black", Value: "Spassky, Boris V."},
							pgn.TagPair{Name: "Result", Value: "1/2-1/2"},
							pgn.TagPair{Name: "a", Value: ""},
							pgn.TagPair{Name: "A", Value: ""},
							pgn.TagPair{Name: "_", Value: ""},
						},
					},
				},
			},
		},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			var unmarshalled pgn.PGN
			err := pgn.Unmarshal(test.in, &unmarshalled)
			if test.errorMessage == "" && err != nil {
				fmt.Println(err)
				t.Fatal("Unexpected error")
			}
			if test.errorMessage != "" && err == nil {
				t.Fatal("Expected error but did not receive one")
			}
			if len(unmarshalled.Games) != len(test.out.Games) {
				fmt.Println("Got:", unmarshalled.Games)
				fmt.Println("Exp:", test.out.Games)
				t.Fatal("Unexpected total games")
			}
			for i, game := range test.out.Games {
				got := unmarshalled.Games[i]
				if len(game.TagPairs) != len(got.TagPairs) {
					fmt.Println("Got:", got.TagPairs)
					fmt.Println("Exp:", game.TagPairs)
					t.Fatal("Unexpected total tag pairs")
				}
			}
		})
	}
}
