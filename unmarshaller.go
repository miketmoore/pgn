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

func Unmarshal(in string, unmarshalled *PGN) error {

	scanner := NewScanner(in)
	lexer := NewLexer(scanner)

	tokens := []Token{}
	err, tokens := lexer.Tokenize(tokens)

	if err != nil {
		return err
	}

	return nil
}
