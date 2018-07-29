package pgn

type PGN struct {
	TagPairs TagPairs
}

type TagPairs struct {
	Event, Site, Date, Round, White, Black, Result string
}

func Parse(s string) PGN {

	return PGN{}
}
