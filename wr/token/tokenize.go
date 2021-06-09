package token

import (
	"regexp"
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
	m, _ := regexp.MatchString(`[\s]`, s)
	return m
}

func isId(s string) bool {
	if s == "EOF" {
		return false
	}
	m, _ := regexp.MatchString(`[A-Za-z_]([A-Z_a-z0-9]+)?`, s)
	return m
}

func isHex(s string) bool {
	if s == "EOF" {
		return false
	}
	m, _ := regexp.MatchString(`([A-Fa-f0-9]+)`, s)
	return m
}

func (c *Chars) Tokenize() []Token {
	tokens := []Token{}
	for {
		for isSpace(currentTok) {
			c.token()
		}

		if currentTok == "\n" {
			Lines++
			c.token()
		}

		if currentTok == "EOF" {
			break
		}

		if isId(currentTok) {
			id := currentTok
			for isId(currentTok) {
				if isId(*c.token()) == false {
					break
				}
				id += currentTok
			}
			tokens = append(tokens, Token{Type: Id, Line: Lines, Text: id})
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
		}

		tokens = append(tokens, Token{Type: Unknown, Line: Lines})
	}
	return tokens
}
