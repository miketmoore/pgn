package pgn

import (
	"bufio"
	"fmt"
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
		if val, ok := parseSevenTagRoster(line, "Event"); ok {
			pgn.TagPairs.Event = val
		} else if val, ok := parseSevenTagRoster(line, "Site"); ok {
			pgn.TagPairs.Site = val
		} else if val, ok := parseSevenTagRoster(line, "Date"); ok {
			pgn.TagPairs.Date = val
		} else if val, ok := parseSevenTagRoster(line, "Round"); ok {
			pgn.TagPairs.Round = val
		} else if val, ok := parseSevenTagRoster(line, "White"); ok {
			pgn.TagPairs.White = val
		} else if val, ok := parseSevenTagRoster(line, "Black"); ok {
			pgn.TagPairs.Black = val
		} else if val, ok := parseSevenTagRoster(line, "Result"); ok {
			pgn.TagPairs.Result = val
		}
	}
	return pgn
}

func parseSevenTagRoster(line, tag string) (string, bool) {
	start := fmt.Sprintf("[%s", tag)
	end := strings.Index(line, "]")
	if strings.Index(line, start) == 0 && end > len(start) {
		sub := line[6:]
		sub = sub[2 : len(sub)-2]
		return sub, true
	}
	return "", false
}
