package pgn

import (
	"errors"
)

type Lexer struct {
	scanner     Scanner
	CurrentRule string
}

func NewLexer(scanner Scanner) Lexer {
	return Lexer{
		scanner:     scanner,
		CurrentRule: "start",
	}
}

type TokenType int

const (
	LB TokenType = iota
	TagNameChar
	TagName
	TagPairOpen
	TagPairClose
	TagValue
	DBLQ
	Letter
	Digit
	SpecialChar
	WS
	RB
	Underscore
	String
	MoveNumber
	File
	Rank
	Piece
)

type Token struct {
	Value    string
	Type     TokenType
	Children []Token
}

/*
digit = "0" ... "9" ;
letter = "A" ... "Z" | "a" ... "z" ;
schar = "!" | '"' | "#" | "$" | "%" | "&" | "'" | "(" | ")"
		| "*" | "+" | "," | "-" | "." | "/" | ":" | ";" | "<" | "="
		| ">" | "?" | "@" | "[" | "\" | "]" | "^" | "_" | "`" |
		| "{" | "|" | "}" | "~" ;
(* Printing character tokens are valid when in ASCII range 32-126 *)
pchar = digit | letter | schar ;
lb = "[" ;
rb = "]" ;
und = "_" ;
tnc = letter | digit | und
tname = tnc , {tnc} ;
dblq = '"' ;
string = dblq , pchar , {pchar} , dblq ;
tpair = lb , tname , string , rb ;
*/
func (l *Lexer) Tokenize(tokens []Token) (error, []Token) {
	l.readWhitespace()

	r := l.scanner.Peek()

	if r == NUL {
		return nil, tokens
	}

	if r == rune('[') {
		// Rule: tpair = lb , tname , string , rb ;

		r := l.scanner.Next()
		tokens = append(tokens, Token{
			Value: "[",
			Type:  TagPairOpen,
		})

		l.readWhitespace()

		value := l.readTagName()
		tokens = append(tokens, Token{
			Value: value,
			Type:  TagName,
		})

		l.readWhitespace()

		err, value := l.readString()
		if err != nil {
			return err, tokens
		}
		tokens = append(tokens, Token{
			Value: value,
			Type:  String,
		})

		l.readWhitespace()

		r = l.scanner.Peek()
		if r != ']' {
			return errors.New("Expected right square bracket but found none"), tokens
		} else {
			l.scanner.Next()
			tokens = append(tokens, Token{
				Value: "]",
				Type:  TagPairClose,
			})
		}

		r = l.scanner.Peek()
		if isNewLine(r) {
			l.scanner.Next()
			return l.Tokenize(tokens)
		}
	} else if isDigit(r) {
		// Rule: movetext = move , {move} ;
		err, movetextTokens := l.readMovetext()
		if err != nil {
			return err, tokens
		}
		for _, t := range movetextTokens {
			tokens = append(tokens, t)
		}
	}
	return nil, tokens
}

// Rule: movetext = move , {move} ;
// Rule: move = move-number , piece , square ;
// Rule: move-number = digit , {digit} , [.] ;
func (l *Lexer) readMovetext() (error, []Token) {
	ok := true
	tokens := []Token{}
	for ok {
		moveNumber := l.readMoveNumber()
		if moveNumber != "" {
			tokens = append(tokens, Token{
				Type:  MoveNumber,
				Value: moveNumber,
			})
		} else {
			ok = false
		}

		err, moveTokens := l.readMove()
		if err != nil {
			return err, tokens
		}
		for _, t := range moveTokens {
			tokens = append(tokens, t)
		}

		l.readWhitespace()

		err, moveTokens = l.readMove()
		if err != nil {
			return err, tokens
		}
		for _, t := range moveTokens {
			tokens = append(tokens, t)
		}

		l.readWhitespace()
	}
	return nil, tokens
}

func (l *Lexer) readMove() (error, []Token) {
	tokens := []Token{}

	// piece is optional, for example e4 indicates that a Pawn (P) moved
	piece := l.readPiece()
	if piece != "" {
		tokens = append(tokens, Token{
			Type:  Piece,
			Value: piece,
		})
	}

	// File is required
	file := l.readFile()
	if file != "" {
		tokens = append(tokens, Token{
			Type:  File,
			Value: file,
		})
	} else if piece != "" {
		return errors.New("File expected to follow piece, but not found."), tokens
	} else {
		// no piece and no file found, so not a move
		return nil, tokens
	}

	rank := l.readRank()
	if rank != "" {
		tokens = append(tokens, Token{
			Type:  Rank,
			Value: rank,
		})
	} else {
		return errors.New("Rank expected to follow file, but not found."), tokens
	}

	return nil, tokens
}

func (l *Lexer) readFile() string {
	r := l.scanner.Peek()
	if isFile(r) {
		l.scanner.Next()
		return string(r)
	}
	return ""
}

func (l *Lexer) readRank() string {
	r := l.scanner.Peek()
	if isRank(r) {
		l.scanner.Next()
		return string(r)
	}
	return ""
}

func (l *Lexer) readMoveNumber() string {
	s := ""

	moveNumber := l.readInteger()
	if moveNumber != "" {
		s = s + moveNumber
	}

	l.skip(isPeriod)
	l.skip(isWhiteSpace)

	return s
}

func (l *Lexer) readPiece() string {

	r := l.scanner.Peek()
	if isPiece(r) {
		l.scanner.Next()
		return string(r)
	}

	return ""
}

func (l *Lexer) readInteger() string {
	ok := true
	s := ""
	for ok {
		peekVal := l.scanner.Peek()
		if !isDigit(peekVal) {
			return s
		}
		l.scanner.Next()
		s = s + string(peekVal)
	}
	return s
}

func (l *Lexer) skip(predicate func(r rune) bool) {
	ok := true
	for ok {
		peekVal := l.scanner.Peek()
		if !predicate(peekVal) {
			return
		}
		l.scanner.Next()
	}
	return
}

func isPeriod(r rune) bool { return r == rune('.') }

// expect one or more tag name characters
// tnc = letter | digit | und
func (l *Lexer) readTagName() string {
	s := ""
	ok := true
	for ok {
		peekVal := l.scanner.Peek()
		if isLetter(peekVal) || isDigit(peekVal) || isUnderscore(peekVal) {
			nextVal := l.scanner.Next()
			s = s + string(nextVal)
		} else {
			ok = false
		}
	}
	return s
}

func (l *Lexer) readWhitespace() {
	ok := true
	for ok {
		peekVal := l.scanner.Peek()
		if !isWhiteSpace(peekVal) {
			return
		}
		l.scanner.Next()
	}
	return
}

func (l *Lexer) readString() (error, string) {
	s := ""
	ok := true

	// check for opening dbl quote
	peekValue := l.scanner.Peek()
	if !isDoubleQuote(peekValue) {
		return errors.New("Expected double quote to denote start of string token"), s
	}
	l.scanner.Next()

	// don't collect the opening quote

	for ok {
		peekVal := l.scanner.Peek()
		if isDoubleQuote(peekVal) {
			l.scanner.Next()
			return nil, s
		} else if isPrintingChar(peekVal) || isWhiteSpace(peekVal) {
			nextVal := l.scanner.Next()
			s = s + string(nextVal)
		} else {
			ok = false
		}
	}
	return nil, s
}

func isLBracket(r rune) bool    { return r == rune('[') }
func isRBracket(r rune) bool    { return r == rune(']') }
func isDoubleQuote(r rune) bool { return r == rune('"') }
func isWhiteSpace(r rune) bool  { return r == rune(' ') }
func isNewLine(r rune) bool     { return r == rune('\n') }
func isLetter(r rune) bool      { return (r >= 65 && r <= 90) || (r >= 97 && r <= 122) }
func isDigit(r rune) bool       { return r >= 48 && r <= 57 }
func isUnderscore(r rune) bool  { return r == rune('_') }

/*
pawn = "P" ;
knight = "N" ;
bishop = "B" ;
rook = "R" ;
queen = "Q" ;
king = "K" ;
*/
func isPiece(r rune) bool { return r == 'P' || r == 'N' || r == 'B' || r == 'R' || r == 'Q' || r == 'K' }
func isFile(r rune) bool  { return r >= 97 && r <= 104 }
func isRank(r rune) bool  { return r >= 49 && r <= 56 }
func isSpecialChar(r rune) bool {
	return (r >= 33 && r <= 47) ||
		(r >= 58 && r <= 64) ||
		(r >= 91 && r <= 96) ||
		(r >= 123 && r <= 126)
}

func isPrintingChar(r rune) bool {
	return isLetter(r) || isDigit(r) || isSpecialChar(r)
}
