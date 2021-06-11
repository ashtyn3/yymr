package parse

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	opcode "yymr/opcodes"
	"yymr/wr/token"
)

var pointer int
var Tokens []token.Token
var Tok token.Token
var Size int = 0

func getTok() {
	if pointer >= len(Tokens) {
		Tok = token.Token{Type: token.End}
		return
	}

	Tok = Tokens[pointer]
	Tok.Line = Tok.Line + 1
	pointer++
}

func peekTok() token.Token {
	if pointer+1 >= len(Tokens) {
		return token.Token{Type: token.End}
	}

	return Tokens[pointer+1]
}

func parseId() ParserToken {
	route := Route{Name: Tok.Text}
	getTok()
	if Tok.Type != token.LeftBrack {
		fmt.Println(strconv.Itoa(Tok.Line) + ": expected '{' after " + route.Name)
		os.Exit(0)
	}

	// set memory address to starting size
	route.Index = Size
	for {
		getTok()
		if Tok.Type == token.RightBrack {
			break
		}
		tok := do(false)
		route.Body = append(route.Body, &tok)
	}

	return ParserToken{Type: "route", Routine: &route, MemIndex: Size}
}

func parseInstruction() ParserToken {
	getTok()
	if Tok.Type != token.Keyword {
		fmt.Println(strconv.Itoa(Tok.Line) + ": unexpected token '" + Tok.Text + "'")
		os.Exit(0)
	}
	instruct := Instruct{Name: Tok.Text}
	rawArgs := []token.Token{}

	for {
		getTok()
		if Tok.Type == token.RightParen {
			break
		}
		if Tok.Text == "," {
			continue
		}
		tok := do(false)
		instruct.Args = append(instruct.Args, &tok)
		rawArgs = append(rawArgs, Tok)
	}

	symbol := KeywordToCode(instruct.Name, rawArgs)
	code := opcode.Opcodes[symbol]
	if code.Size == 0x0000 {
		fmt.Println(strconv.Itoa(Tok.Line) + ": uknown instruction '" + Tok.Text + "'")
		os.Exit(0)
	}

	if code.Size != len(instruct.Args)+1 {
		fmt.Println(strconv.Itoa(Tok.Line) + ": too many arguements for " + instruct.Name + " instruction")
		os.Exit(0)
	}

	instruct.Code = code.Code
	Size += code.Size

	return ParserToken{Type: "instruction", Instruction: &instruct, MemIndex: Size}
}

func parseLit() ParserToken {
	lit := Lit{}

	if Tok.Type == token.Hex {
		lit.Type = "hex"
		num, _ := strconv.ParseUint(Tok.Text, 16, 16)
		lit.Hex.Value = uint16(num)
	}

	return ParserToken{Type: "lit", Literal: &lit}
}

func parseRef() ParserToken {
	ref := Refer{}

	if Tok.Type == token.RegisterId {
		ref.SubType = "reg"
		if strings.HasPrefix(strings.ToLower(Tok.Text), "r") == false {
			fmt.Println(strconv.Itoa(Tok.Line) + ": unknown register " + Tok.Text)
			os.Exit(0)
		}
		ref.Id = strings.TrimPrefix(Tok.Text, "r")
	}

	return ParserToken{Type: "ref", Reference: &ref}
}

func do(move bool) ParserToken {
	if move {
		getTok()
	}
	if Tok.Type == token.Id {
		return parseId()
	} else if Tok.Type == token.LeftParen {
		return parseInstruction()
	} else if Tok.Type == token.Hex {
		return parseLit()
	} else if Tok.Type == token.RegisterId {
		return parseRef()
	}
	return ParserToken{}
}

func Parse(toks []token.Token) []ParserToken {
	Tokens = toks
	parserToks := []ParserToken{}
	for {
		end := do(true)
		if end.Type == "" {
			break
		}
		parserToks = append(parserToks, end)
	}
	return parserToks
}
