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

func (p PGN) String() string {
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
		str += fmt.Sprintf("%d. %s %s", i+1, entry.White, entry.Black)
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
	White, Black string
	Comments     []Comment
}

func Parse(raw string) PGN {
	r := strings.NewReader(raw)
	scanner := bufio.NewScanner(r)
	pgn := PGN{}
	section := "tagpair"
	movetextLines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if section == "tagpair" {
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
				section = "movetext"
			}
		} else if section == "movetext" {
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
				White:    moves[0],
				Black:    moves[1],
				Comments: comments,
			})
		}

	}

	return mt
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
