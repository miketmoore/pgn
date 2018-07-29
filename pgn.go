package pgn

import (
	"bufio"
	"regexp"
	"strings"
)

type PGN struct {
	TagPairs TagPairs
	Movetext Movetext
}

type TagPairs struct {
	Event, Site, Date, Round, White, Black, Result string
}

type Movetext [][]string

func Parse(raw string) PGN {
	r := strings.NewReader(raw)
	scanner := bufio.NewScanner(r)
	pgn := PGN{}
	section := "tagpair"
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

		}

	}
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
