package pgn

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type PGN struct {
	TagPairs TagPairs
	Movetext Movetext
}

// func (p *PGN) RecordWhiteMove(file, rank string, capture bool, piece string, check bool) {
// 	s := buildMoveString(file, rank, capture, piece, check)
// 	p.Movetext = append(p.Movetext, MovetextEntry{White: s})
// }

// func (p *PGN) RecordBlackMove(file, rank string, capture bool, piece string, check bool) {
// 	s := buildMoveString(file, rank, capture, piece, check)
// 	move := p.Movetext[len(p.Movetext)-1]
// 	move.Black = s
// 	p.Movetext[len(p.Movetext)-1] = move
// }

func buildMoveString(file, rank string, capture bool, piece string, check bool) string {
	var s string
	if capture {
		if piece != "pawn" {
			s = fmt.Sprintf("%sx%s%s", getPieceInitialByFullname(piece), file, rank)
		} else {
			s = fmt.Sprintf("x%s%s", file, rank)
		}
	} else {
		if piece != "pawn" {
			s = fmt.Sprintf("%s%s%s", getPieceInitialByFullname(piece), file, rank)
		} else {
			s = fmt.Sprintf("%s%s", file, rank)
		}

	}
	if check {
		s += "+"
	}
	return s
}

func getPieceInitialByFullname(fullname string) string {
	switch fullname {
	case "pawn":
		return "P"
	case "knight":
		return "N"
	case "bishop":
		return "B"
	case "rook":
		return "R"
	case "queen":
		return "Q"
	case "king":
		return "K"
	}
	return ""
}

func (p *PGN) String() string {
	str := ""
	str += fmt.Sprintf("[Event \"%s\"]\n", p.TagPairs.Event)
	str += fmt.Sprintf("[Site \"%s\"]\n", p.TagPairs.Site)
	str += fmt.Sprintf("[Date \"%s\"]\n", p.TagPairs.Date)
	str += fmt.Sprintf("[Round \"%s\"]\n", p.TagPairs.Round)
	str += fmt.Sprintf("[White \"%s\"]\n", p.TagPairs.White)
	str += fmt.Sprintf("[Black \"%s\"]\n", p.TagPairs.Black)
	str += fmt.Sprintf("[Result \"%s\"]\n", p.TagPairs.Result)
	str += "\n"
	for i, entry := range p.Movetext {
		str += fmt.Sprintf("%d. %s %s", i+1, entry.White.Original, entry.Black.Original)
		if len(entry.Comments) > 0 {
			for _, comment := range entry.Comments {
				str += fmt.Sprintf(" {%s}", comment)
			}
		}
		if i < len(p.Movetext)-1 {
			str += " "
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

type Move struct {
	Original     string
	File         File
	Rank         Rank
	Piece        Piece
	Capture      bool
	Check        bool
	Disambiguate Disambiguate
}

type Section string

const (
	SectionTagPair  Section = "tagpair"
	SectionMovetext Section = "movetext"
)

func Unmarshal(raw string) PGN {
	r := strings.NewReader(raw)
	scanner := bufio.NewScanner(r)
	pgn := PGN{}
	section := SectionTagPair
	movetextLines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if section == SectionTagPair {
			key, val, ok := parseSevenTagRoster(line)
			if ok {
				switch key {
				case "Event":
					pgn.TagPairs.Event = val
				case "Site":
					pgn.TagPairs.Site = val
				case "Date":
					pgn.TagPairs.Date = val
				case "Round":
					pgn.TagPairs.Round = val
				case "White":
					pgn.TagPairs.White = val
				case "Black":
					pgn.TagPairs.Black = val
				case "Result":
					pgn.TagPairs.Result = val
				}
			}
			if strings.TrimSpace(line) == "" {
				section = SectionMovetext
			}
		} else if section == SectionMovetext {
			movetextLines = append(movetextLines, line)
		}

	}
	pgn.Movetext = parseMovetext(movetextLines)
	return pgn
}

func parseSevenTagRoster(line string) (string, string, bool) {
	var re = regexp.MustCompile(`\[(.*) "(.*)"\]`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 3 {
		return matches[1], matches[2], true
	}
	return "", "", false
}

func parseMovetext(lines []string) Movetext {
	mt := Movetext{}
	str := strings.Join(lines, " ")

	re := regexp.MustCompile(`[1-9].?\. `)
	split := re.Split(str, -1)

	for _, movetextSection := range split {
		if strings.TrimSpace(movetextSection) != "" {
			comments := parseComments(movetextSection)
			moves := strings.Split(strings.TrimSpace(movetextSection), " ")
			for i := 0; i < len(moves); i++ {
				moves[i] = strings.TrimSpace(moves[i])
			}
			mt = append(mt, MovetextEntry{
				White:    parseMove(moves[0]),
				Black:    parseMove(moves[1]),
				Comments: comments,
			})
		}

	}

	return mt
}

var nonPawnMove = regexp.MustCompile(`^([PNBRQK])([a-h])([1-8])$`)
var disambiguateByFileRe = regexp.MustCompile(`^([PNBRQK])([a-h])([a-h])([1-8])$`)

func parseMove(move string) Move {
	m := Move{Original: move}
	if move == "O-O" {
		m.Piece = PieceKing
	} else if move == "O-O-O" {
		m.Piece = PieceKing
	} else if len(move) == 2 {
		m.Piece = PiecePawn
		m.File = File(move[0])
		m.Rank = Rank(move[1])
	} else if len(move) == 3 {
		matches := nonPawnMove.FindStringSubmatch(move)
		if len(matches[1:]) == 3 {
			m.Piece = Piece(move[0])
			m.File = File(move[1])
			m.Rank = Rank(move[2])
		}
	} else if len(move) == 4 {
		byFileMatches := disambiguateByFileRe.FindStringSubmatch(move)
		if strings.Index(move, "+") == 3 {
			m.Piece = Piece(move[0])
			m.File = File(move[1])
			m.Rank = Rank(move[2])
			m.Check = true
		} else if strings.Index(move, "x") == 1 {
			if regexp.MustCompile(`^[a-h]{1}`).MatchString(move) {
				m.Piece = PiecePawn
				m.Disambiguate = Disambiguate{
					File: File(move[0]),
				}
				m.File = File(move[2])
				m.Rank = Rank(move[3])
			} else {
				m.Piece = Piece(move[0])
				m.File = File(move[2])
				m.Rank = Rank(move[3])
			}
			m.Capture = true
		} else if len(byFileMatches) > 0 && len(byFileMatches[1:]) == 4 {
			m.Piece = Piece(move[0])
			m.Disambiguate = Disambiguate{
				File: File(move[1]),
			}
			m.File = File(move[2])
			m.Rank = Rank(move[3])
		}
	} else if len(move) == 5 {
		if regexp.MustCompile(`^[BNRQK]x[a-h][1-8]\+`).MatchString(move) {
			m.Piece = Piece(move[0])
			m.Capture = true
			m.File = File(move[2])
			m.Rank = Rank(move[3])
			m.Check = true
		} else {
			panic(move)
		}
	}
	return m
}

func parseComments(val string) []Comment {
	commentsRe := regexp.MustCompile(`\{(.*)\}`)
	matches := commentsRe.FindAllStringSubmatch(val, -1)
	comments := []Comment{}
	if len(matches) > 0 {
		for _, match := range matches {
			comments = append(comments, Comment(match[1]))
		}
	}
	return comments
}
