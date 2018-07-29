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

type MovetextEntry struct {
	White, Black string
	Comments     []string
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
		// fmt.Println(matches[1])
		// fmt.Println(matches[2])
		return matches[1], matches[2], true
	}
	return "", "", false
}

func parseMovetext(lines []string) Movetext {
	mt := Movetext{}
	str := strings.Join(lines, " ")

	re := regexp.MustCompile(`[1-9]\. `)
	split := re.Split(str, -1)

	for _, val := range split {
		if strings.TrimSpace(val) != "" {
			// TODO get comments
			moves := strings.Split(strings.TrimSpace(val), " ")
			for i := 0; i < len(moves); i++ {
				moves[i] = strings.TrimSpace(moves[i])
				fmt.Println(moves[i])
			}
			fmt.Println(moves)
			mt = append(mt, MovetextEntry{
				White: moves[0],
				Black: moves[1],
			})
		}

	}

	return mt
}
