package main

import (
	"fmt"
	"os"
	"strings"
	"yymr/wr/assemble"
	"yymr/wr/parse"
	"yymr/wr/token"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("WR", "The tokenizer and parser for yymr-WR (written repersentation)")

	f := parser.File("i", "input", os.O_RDWR, 0600, &argparse.Options{Required: true, Help: "Specifies input file"})
	output := parser.String("o", "output", &argparse.Options{Default: "a.out"})

	parser.Parse(os.Args)

	if *f == *new(os.File) {
		fmt.Println(parser.Help("Needs filename"))
		os.Exit(0)
	}

	defer f.Close()

	info, _ := f.Stat()

	file := make([]byte, info.Size())

	f.Read(file)

	var buf token.Chars = strings.Split(string(file), "")
	buf = append(buf, "EOF")

	toks := parse.Parse(buf.Tokenize())
	b := assemble.Build{Filename: *output, Toks: toks}
	b.Assemble()
}
