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
				if game != got {
					t.Fatal("Unexpected game")
				}
			}
		})
	}
}
