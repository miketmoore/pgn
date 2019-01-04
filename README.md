# :construction: PGN Parser

This is a Portable Game Notation (PGN) parser written in Go. Note: this is a work in progress and is the first parser I've built, so it is a learning project.

## Development

### Get

```
go get github.com/miketmoore/pgn
```

### Install & Run

```
cd ~/go/src/github.com/miketmoore/pgn/
go install cmd/parser/parser.go
pgn -input ./data/games/fischer_spassky_1992_11_04.pgn
```

### Tests

```
go test
```

## Production Rules

List of production rules in EBNF grammar.

```
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
(/* Tag Name Character */)
tnc = letter | digit | und
tname = tnc , {tnc} ;
dblq = '"' ;
string = dblq , pchar , {pchar} , dblq ;
tpair = lb , tname , string , rb ;

file = "a" ... "h" ;
rank = "1" ... "8" ;
square = file , rank ;

pawn = "P" ;
knight = "N" ;
bishop = "B" ;
rook = "R" ;
queen = "Q" ;
king = "K" ;
piece = pawn | knight | bishop | rook | queen | king ;

move-number = digit , {digit} , [.] ;
move = move-number , piece , square ;

capture-move = "x" , square ;
pawn-capture-move = file , "x" , square ;
castle-kingside	= "O-O" ;
castle-queenside = "O-O-O" ;
promotion-piece	= knight | bishop | rook | queen ;
pawn-promotion = square , "=" , promotion-piece ;
checking-move = move , "+" ;
checkmating-move = move , "#" ;

movetext = move , {move} ;
```

## References

- https://opensource.apple.com/source/Chess/Chess-110.0.6/Documentation/PGN-Standard.txt
- http://www.saremba.de/chessgml/standards/pgn/pgn-complete.htm
