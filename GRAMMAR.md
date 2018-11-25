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
