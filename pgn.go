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
		// start := strings.Index(line, "[")
		end := strings.Index(line, "]")
		if strings.Index(line, "[Event") == 0 && end > len("[Event") {
			sub := line[6:]
			fmt.Println(sub)
			sub = sub[2 : len(sub)-2]
			fmt.Println(sub)
			pgn.TagPairs.Event = sub
		}
	}
	return pgn
}
