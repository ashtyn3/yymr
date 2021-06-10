package main

import (
	"fmt"
	"os"
	"strings"
	"yymr/wr/token"
)

func main() {

	file, _ := os.ReadFile("hi.s")

	var buf token.Chars = strings.Split(string(file), "")
	buf = append(buf, "EOF")

	fmt.Println(buf.Tokenize())
	os.Exit(0)
}
