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
			out:  pgn.PGN{},
		},
		{
			name: "Tag pair",
			in:   "",
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
				t.Fatal("Unexpected total games")
			}
			for i, game := range test.out.Games {
				got := unmarshalled.Games[i]
				if len(game.TagPairs) != len(got.TagPairs) {
					t.Fatal("Unexpected total tag pairs")
				}
			}
		})
	}
}
