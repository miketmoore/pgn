package pgn

import (
	"errors"
	"strconv"
)

type Game struct {
	TagPairs []TagPair
	Movetext []Movetext
}

type TagPair struct {
	Name, Value string
}

type Movetext struct {
	Num   int
	White Move
	Black Move
}

type File string
type Rank int

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
	Rank1 Rank = 1
	Rank2 Rank = 2
	Rank3 Rank = 3
	Rank4 Rank = 4
	Rank5 Rank = 5
	Rank6 Rank = 6
	Rank7 Rank = 7
	Rank8 Rank = 8
)

type Move struct {
	File File
	Rank Rank
}

type PGN struct {
	Games []Game
}

type unmarshaller struct {
	tokens []Token
	index  int
}

func Unmarshal(in string, unmarshalled *PGN) error {

	scanner := NewScanner(in)
	lexer := NewLexer(scanner)

	err, tokens := lexer.Tokenize()

	if err != nil {
		return err
	}

	u := unmarshaller{tokens: tokens}

	game := Game{}

	ok := true
	for ok {
		tagPair := u.readTagPair()
		if tagPair != nil {
			game.TagPairs = append(game.TagPairs, *tagPair)
		} else {
			ok = false
		}
	}

	// move text
	ok = true
	for ok {
		token := u.peek()
		if token != nil && token.Type == TokenMoveNumber {
			u.next()

			// convert value to int
			i, err := strconv.Atoi(token.Value)
			if err != nil {
				return err
			}

			// create movetext instance
			m := Movetext{
				Num: i,
			}

			// parse white move
			token = u.peek()
			if token == nil {
				return errors.New("white movetext token is nil")
			}
			if token.Type == TokenFile {
				m.White = Move{File: File(token.Value)}
			}

			u.next()

			// parse black move
			token = u.peek()
			if token == nil {
				return errors.New("black movetext token is nil")
			}
			if token.Type == TokenFile {
				m.Black = Move{File: File(token.Value)}
			}

			u.next()

			game.Movetext = append(game.Movetext, m)
		} else {
			ok = false
		}
	}

	unmarshalled.Games = append(unmarshalled.Games, game)

	return nil
}

func (u *unmarshaller) readTagPair() *TagPair {
	tagPair := TagPair{}

	token := u.peek()
	if token != nil && token.Type == TokenTagName {
		u.next()
		tagPair.Name = token.Value
		token = u.peek()
		if token != nil && token.Type == TokenTagValue {
			u.next()
			tagPair.Value = token.Value
			return &tagPair
		}
	}

	return nil
}

func (u *unmarshaller) readEOF() bool {
	if u.peek().Type == TokenEOF {
		u.next()
		return true
	}
	return false
}

func (u *unmarshaller) peek() *Token {
	for _, t := range u.tokens {
		return &t
	}
	return nil
}

func (u *unmarshaller) next() *Token {
	for _, t := range u.tokens {
		u.tokens = u.tokens[1:]
		return &t
	}
	return nil
}
