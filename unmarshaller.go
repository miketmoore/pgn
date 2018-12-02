package pgn

type Game struct {
	TagPairs []TagPair
}

type TagPair struct {
	Name, Value string
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

	unmarshalled.Games = append(unmarshalled.Games, game)

	return nil
}

func (u *unmarshaller) readTagPair() *TagPair {
	tagPair := TagPair{}

	token := u.peek()
	if token.Type == TokenTagName {
		u.next()
		tagPair.Name = token.Value
		token = u.peek()
		if token.Type == TokenTagValue {
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
