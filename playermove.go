package pgn

type PlayerMove struct {
	File           File
	Rank           Rank
	CastleKingside bool
}
