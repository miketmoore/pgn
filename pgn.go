package pgn

import (
	"fmt"
)

// PGN represents an unmarshalled PGN format string
// It represents zero or more games
type PGN struct {
	Games []Game
}

type Game struct {
	TagPairs TagPairs
	Movetext Movetext
}

// String returns a PGN formatted string
func (p *PGN) String() string {
	str := ""
	for _, game := range p.Games {
		str += fmt.Sprintf("[Event \"%s\"]\n", game.TagPairs.Event)
		str += fmt.Sprintf("[Site \"%s\"]\n", game.TagPairs.Site)
		str += fmt.Sprintf("[Date \"%s\"]\n", game.TagPairs.Date)
		str += fmt.Sprintf("[Round \"%s\"]\n", game.TagPairs.Round)
		str += fmt.Sprintf("[White \"%s\"]\n", game.TagPairs.White)
		str += fmt.Sprintf("[Black \"%s\"]\n", game.TagPairs.Black)
		str += fmt.Sprintf("[Result \"%s\"]\n", game.TagPairs.Result)
		str += "\n"
		for i, entry := range game.Movetext {
			str += fmt.Sprintf("%d. %s %s", i+1, entry.White.Original, entry.Black.Original)
			if len(entry.Comments) > 0 {
				for _, comment := range entry.Comments {
					str += fmt.Sprintf(" {%s}", comment)
				}
			}
			if i < len(game.Movetext)-1 {
				str += " "
			}
		}
	}
	return str
}

type TagPairs struct {
	Event, Site, Date, Round, White, Black, Result string
}

type Movetext []MovetextEntry

type Comment string

type MovetextEntry struct {
	White, Black Move
	Comments     []Comment
}

type File string
type Rank string

const (
	FileA File = "a"
	FileB File = "b"
	FileC File = "c"
	FileD File = "d"
	FileE File = "e"
	FileF File = "f"
	FileG File = "g"
	FileH File = "h"
)

const (
	Rank1 Rank = "1"
	Rank2 Rank = "2"
	Rank3 Rank = "3"
	Rank4 Rank = "4"
	Rank5 Rank = "5"
	Rank6 Rank = "6"
	Rank7 Rank = "7"
	Rank8 Rank = "8"
)

type Piece string

const (
	PiecePawn   Piece = "P"
	PieceBishop Piece = "B"
	PieceRook   Piece = "R"
	PieceKnight Piece = "N"
	PieceQueen  Piece = "Q"
	PieceKing   Piece = "K"
)

// Disambiguate is used to notate that the PGN movetext is ambiguous
// Note that the above disambiguation is needed only to distinguish among moves of
// the same piece type to the same square;
type Disambiguate struct {
	File File
	Rank Rank
}

// Move represents a white or a black move
type Move struct {
	Original     string
	File         File
	Rank         Rank
	Piece        Piece
	Capture      bool
	Check        bool
	Disambiguate Disambiguate
}

// Section represents a section of the PGN formatted string
type Section string

const (
	SectionTagPair  Section = "tagpair"
	SectionMovetext Section = "movetext"
)
