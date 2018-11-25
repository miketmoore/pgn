package pgn

// Notes
// https://blog.golang.org/strings
// In Go, a string is a read-only and arbitrary slice of bytes.
// It is arbitrary because it is not constrained to hold only a specific format,
// such as Unicode or UTF-8.
// Keep in mind that bytes do not equal characters.
// A rune is a code point with an int32 value.

// Scanner scans a string rune by rune and exposes a simple API for moving
// through the stream
type Scanner struct {
	stream string
	index  int
}

// NewScanner returns an instance of Scanner
func NewScanner(in string) Scanner {
	return Scanner{
		stream: in,
	}
}

// Peek returns the next rune without discarding the current
func (s *Scanner) Peek() rune {
	// A "for range" loop is used because it decodes one UTF-8-encoded rune on
	// each iteration.
	for _, r := range s.stream {
		return r
	}
	return NUL
}

func (s *Scanner) Next() rune {
	for _, r := range s.stream {
		s.stream = s.stream[1:]
		return r
	}
	return NUL
}

const NUL = rune(0)
