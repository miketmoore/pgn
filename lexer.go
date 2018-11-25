package pgn

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
)

type Token struct {
	Value    string
	Type     TokenType
	Children []Token
}

func NewToken(t TokenType, v string) Token {
	return Token{Type: t, Value: v}
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
func (l *Lexer) Tokenize() (bool, []Token) {
	tokens := []Token{}

	ok, r := l.scanner.Next()
	if !ok {
		return false, tokens
	}
	if r == rune('[') {
		// tpair = lb , tname , string , rb ;
		ok, value := l.readTagName()
		if !ok {
			return false, tokens
		}
		tokens = append(tokens, Token{
			Value: value,
			Type:  TagName,
		})
	}
	return true, tokens
}

// expect one or more tag name characters
// tnc = letter | digit | und
func (l *Lexer) readTagName() (bool, string) {
	s := ""
	ok := true
	i := 0
	for ok {
		peekOk, peekVal := l.scanner.Peek()
		if !peekOk {
			return false, s
		}
		if isLetter(peekVal) || isDigit(peekVal) || isUnderscore(peekVal) {
			nextOk, nextVal := l.scanner.Next()
			if !nextOk {
				return false, s
			}
			s = s + string(nextVal)
		} else {
			ok = false
		}
		i++
	}
	return true, s
}

func isLBracket(r rune) bool    { return r == rune('[') }
func isRBracket(r rune) bool    { return r == rune(']') }
func isDoubleQuote(r rune) bool { return r == rune('"') }
func isWhiteSpace(r rune) bool  { return r == rune(' ') }
func isNewLine(r rune) bool     { return r == rune('\n') }
func isLetter(r rune) bool      { return (r >= 65 && r <= 90) || (r >= 97 && r <= 122) }
func isDigit(r rune) bool       { return r >= 48 && r <= 57 }
func isUnderscore(r rune) bool  { return r == rune('_') }

func isSpecialChar(r rune) bool {
	return (r >= 33 && r <= 47) ||
		(r >= 58 && r <= 64) ||
		(r >= 91 && r <= 96) ||
		(r >= 123 && r <= 126)
}

func isPrintingChar(r rune) bool {
	return isLetter(r) || isDigit(r) || isSpecialChar(r)
}
