package main

import (
	"os"
	"strings"
	"yymr/wr/assemble"
	"yymr/wr/parse"
	"yymr/wr/token"
)

func main() {

	file, _ := os.ReadFile("hi.s")

	var buf token.Chars = strings.Split(string(file), "")
	buf = append(buf, "EOF")

	toks := parse.Parse(buf.Tokenize())
	b := assemble.Build{Filename: "a.out", Exec: true, Toks: toks}
	b.Assemble()
}
