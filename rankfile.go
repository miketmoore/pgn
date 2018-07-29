package pgn

// Rank is a custom type that represents a horizontal row (rank) on the chess board
type Rank int

// File is a custom type that represents a vertical column (file) on the chess board
type File int

const (
	RankNone Rank = 0
	Rank1    Rank = 1
	Rank2    Rank = 2
	Rank3    Rank = 3
	Rank4    Rank = 4
	Rank5    Rank = 5
	Rank6    Rank = 6
	Rank7    Rank = 7
	Rank8    Rank = 8

	FileNone File = 0
	FileA    File = 1
	FileB    File = 2
	FileC    File = 3
	FileD    File = 4
	FileE    File = 5
	FileF    File = 6
	FileG    File = 7
	FileH    File = 8
)

var FileView = map[File]string{
	FileA: "a",
	FileB: "b",
	FileC: "c",
	FileD: "d",
	FileE: "e",
	FileF: "f",
	FileG: "g",
	FileH: "h",
}

var RankView = map[Rank]string{
	Rank1: "1",
	Rank2: "2",
	Rank3: "3",
	Rank4: "4",
	Rank5: "5",
	Rank6: "6",
	Rank7: "7",
	Rank8: "8",
}
