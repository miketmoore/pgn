package pgn

import (
	"bufio"
	"regexp"
	"strings"
)

// Unmarshal parses a string PGN format to a PGN struct
func Unmarshal(raw string, data *PGN) error {
	if strings.TrimSpace(raw) == "" {
		return UnmarshalError{
			Value: "",
		}
	}
	r := strings.NewReader(raw)
	scanner := bufio.NewScanner(r)
	section := SectionTagPair
	movetextLines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if section == SectionTagPair {
			key, val, ok := unmarshalSevenTagRoster(line)
			if ok {
				switch key {
				case "Event":
					data.TagPairs.Event = val
				case "Site":
					data.TagPairs.Site = val
				case "Date":
					data.TagPairs.Date = val
				case "Round":
					data.TagPairs.Round = val
				case "White":
					data.TagPairs.White = val
				case "Black":
					data.TagPairs.Black = val
				case "Result":
					data.TagPairs.Result = val
				}
			}
			if strings.TrimSpace(line) == "" {
				section = SectionMovetext
			}
		} else if section == SectionMovetext {
			movetextLines = append(movetextLines, line)
		}

	}
	data.Movetext = unmarshalMovetext(movetextLines)
	return nil
}

func unmarshalSevenTagRoster(line string) (string, string, bool) {
	var re = regexp.MustCompile(`\[(.*) "(.*)"\]`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 3 {
		return matches[1], matches[2], true
	}
	return "", "", false
}

func unmarshalMovetext(lines []string) Movetext {
	mt := Movetext{}
	str := strings.Join(lines, " ")

	re := regexp.MustCompile(`[1-9].?\. `)
	split := re.Split(str, -1)

	for _, movetextSection := range split {
		if strings.TrimSpace(movetextSection) != "" {
			comments := unmarshalComments(movetextSection)
			moves := strings.Split(strings.TrimSpace(movetextSection), " ")
			for i := 0; i < len(moves); i++ {
				moves[i] = strings.TrimSpace(moves[i])
			}
			mt = append(mt, MovetextEntry{
				White:    unmarshalMove(moves[0]),
				Black:    unmarshalMove(moves[1]),
				Comments: comments,
			})
		}

	}

	return mt
}

var nonPawnMove = regexp.MustCompile(`^([PNBRQK])([a-h])([1-8])$`)
var disambiguateByFileRe = regexp.MustCompile(`^([PNBRQK])([a-h])([a-h])([1-8])$`)

func unmarshalMove(move string) Move {
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

func unmarshalComments(val string) []Comment {
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
