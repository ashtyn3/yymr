package main

import (
	"fmt"
	"strings"
	"yymr/wr/token"
)

func main() {

	for {
		var f token.Chars
		var l []byte
		fmt.Scanln(&l)

		f = strings.Split(string(l), "")
		f = append(f, "EOF")

		fmt.Println(f.Tokenize())
	}
}
