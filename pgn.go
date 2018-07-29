package pgn

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type PGN struct {
	TagPairs TagPairs
}

type TagPairs struct {
	Event, Site, Date, Round, White, Black, Result string
}

func Parse(raw string) PGN {
	r := strings.NewReader(raw)
	scanner := bufio.NewScanner(r)
	pgn := PGN{}
	for scanner.Scan() {
		line := scanner.Text()
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
	}
	return pgn
}

func parseSevenTagRoster(line string) (string, string, bool) {
	var re = regexp.MustCompile(`\[(.*) "(.*)"\]`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 3 {
		fmt.Println(matches[1])
		fmt.Println(matches[2])
		return matches[1], matches[2], true
	}
	return "", "", false
}
