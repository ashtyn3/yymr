package main

import (
	"os"
	"strings"
	"yymr/wr/parse"
	"yymr/wr/token"
)

func main() {

	file, _ := os.ReadFile("hi.s")

	var buf token.Chars = strings.Split(string(file), "")
	buf = append(buf, "EOF")

	parse.Parse(buf.Tokenize())
	os.Exit(0)
}
