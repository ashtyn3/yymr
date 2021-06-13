package token

import (
	"regexp"
	"strings"
)

type Chars []string

var Pointer int = -1
var Lines int = 0

var currentTok string = " "
var END string = "EOF"

func (c Chars) peekToken() *string {
	if Pointer+1 >= len(c) {
		return &END
	}
	return &c[Pointer+1]
}

func (c Chars) token() *string {
	Pointer++
	if Pointer >= len(c) {
		return &END
	}
	currentTok = c[Pointer]
	return &currentTok
}

func isSpace(s string) bool {
	m, _ := regexp.MatchString(`[\s\r\t]`, s)
	return m
}

func isIdStart(s string) bool {
	if s == "EOF" {
		return false
	}
	m, _ := regexp.MatchString(`[A-Za-z_]`, s)
	return m
}

func isId(s string) bool {
	if s == "EOF" {
		return false
	}
	m, _ := regexp.MatchString(`[A-Z_a-z0-9]+`, s)
	return m
}

func isHex(s string) bool {
	if s == "EOF" {
		return false
	}
	m, _ := regexp.MatchString(`([A-Fa-f0-9]+)`, s)
	return m
}

func (c Chars) Tokenize() []Token {
	tokens := []Token{}
	for Pointer < len(c) {
		if currentTok == "\n" {
			Lines++
			c.token()
			continue
		}

		for isSpace(currentTok) {
			c.token()
		}

		if strings.TrimSpace(currentTok) == "EOF" || Pointer == len(c)-1 {
			break
		}

		if isIdStart(currentTok) {
			id := currentTok
			for isId(currentTok) {
				if isId(*c.token()) == false {
					break
				}
				id += currentTok
			}
			if id == "mov" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			} else if id == "push" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			} else if id == "pop" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			} else if id == "call" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			} else if id == "add" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			} else if id == "mul" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			} else if id == "div" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			} else if id == "inc" || id == "dec" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			} else if id == "jmpE" || id == "jmpNE" || id == "jmpL" || id == "jmpLE" || id == "jmpG" || id == "jmpGE" {
				tokens = append(tokens, Token{Type: Keyword, Line: Lines, Text: id})
				continue
			}

			tokens = append(tokens, Token{Type: Id, Line: Lines, Text: id})

			continue
		}

		if currentTok == "0" && *c.peekToken() == "x" {
			c.token()
			c.token()
			id := currentTok
			for isHex(currentTok) {
				if isHex(*c.token()) == false {
					break
				}
				id += currentTok
			}
			tokens = append(tokens, Token{Type: Hex, Line: Lines, Text: id})

			continue
		}

		if currentTok == "@" {
			c.token()
			id := currentTok
			for isId(currentTok) {
				if isId(*c.token()) == false {
					break
				}
				id += currentTok
			}
			tokens = append(tokens, Token{Type: RegisterId, Line: Lines, Text: id})
			continue
		}

		if currentTok == "$" {
			c.token()
			id := currentTok
			for isId(currentTok) {
				if isId(*c.token()) == false {
					break
				}
				id += currentTok
			}
			tokens = append(tokens, Token{Type: RouteId, Line: Lines, Text: id})

			continue
		}

		if currentTok == "%" {
			c.token()
			id := currentTok
			for isId(currentTok) {
				if isId(*c.token()) == false {
					break
				}
				id += currentTok
			}
			tokens = append(tokens, Token{Type: MemId, Line: Lines, Text: id})

			continue
		}

		if currentTok == "\"" {
			for {
				c.token()
				if currentTok == "\n" || currentTok == "EOF" {
					break
				}
			}
			continue
		}

		if currentTok == ";" && *c.peekToken() == ";" {
			c.token()
			tokens = append(tokens, Token{Type: LineSep, Line: Lines, Text: ";;"})
			c.token()
			continue

		}

		if currentTok == "," {
			tokens = append(tokens, Token{Type: Unknown, Line: Lines, Text: ","})
			c.token()
			continue
		}
		if currentTok == "(" {
			tokens = append(tokens, Token{Type: LeftParen, Line: Lines, Text: "("})
			c.token()
			continue
		}

		if currentTok == ")" {
			tokens = append(tokens, Token{Type: RightParen, Line: Lines, Text: ")"})
			c.token()
			continue
		}

		if currentTok == "{" {
			tokens = append(tokens, Token{Type: LeftBrack, Line: Lines, Text: "{"})
			c.token()
			continue
		}

		if currentTok == "}" {
			tokens = append(tokens, Token{Type: RightBrack, Line: Lines, Text: "}"})
			c.token()
			continue
		}
		tokens = append(tokens, Token{Type: Unknown, Line: Lines, Text: currentTok})
	}
	return tokens
}
