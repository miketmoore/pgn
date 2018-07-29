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
	fmt.Println(str)

	// split by numbers /[0-1]\.(.*)/
	// x := strings.Split(str, ". ")

	// for _, val := range x {
	// 	fmt.Println(val)
	// }

	// re := regexp.MustCompile(`([0-1]\. )(.*)`)
	// split := re.Split(str, -1)
	// for _, val := range split {
	// 	fmt.Println(val)
	// }

	// str := "1. e4 e5 2. Nf3 Nc6 3. Bb5 a6"
	re := regexp.MustCompile(`[1-9]\. `)
	split := re.Split(str, -1)
	// final := [][]string{}
	for _, val := range split {
		if strings.TrimSpace(val) != "" {
			moves := strings.Split(strings.TrimSpace(val), " ")
			for i := 0; i < len(moves); i++ {
				moves[i] = strings.TrimSpace(moves[i])
			}
			fmt.Println(moves)

			// final = append(final, moves)
			// mt = append(mt, moves)
			mt = append(mt, MovetextEntry{
				White: moves[0],
				Black: moves[1],
			})
		}

	}

	return mt
}
