## Production Rules

List of production rules (tokens) described in EBNF grammar.

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

## Pseudo Tokenization

```
IsTagPair
    IsLeftBracket => IsTagName => IsString => IsRightBracket => IsLineFeed
IsTagName
    IsTagNameChar {IsTagNameChar}
IsTagNameChar
    IsLetter | IsDigit | IsUnderscore
IsString
    IsDBLQ => IsPrintingChar {IsPrintingChar} => IsDBLQ
IsPrintingChar
    IsDigit | IsLetter | IsSpecialChar
```

## Parse Tree

```
TPAIR
  |
  ---> LB , TNAME , STRING , RB ;
              |        |
              |        ---> DBLQ , PCHAR , {PCHAR} , DBLQ
              |                      |
              |                      ---> DIGIT | LETTER | SCHAR
              |
              ---> TNAMECHAR , {TNAMECHAR}
                       |
                       ---> LETTER | DIGIT | UND
```
