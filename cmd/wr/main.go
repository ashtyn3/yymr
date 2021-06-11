package main

import (
	"os"
	"strings"
	"yymr/wr/assemble"
	"yymr/wr/parse"
	"yymr/wr/token"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("WR", "The tokenizer and parser for yymr-WR (written repersentation)")

	f := parser.File("i", "input", 0060, 0060, &argparse.Options{Required: true})
	output := parser.String("o", "output", &argparse.Options{Default: "a.out"})

	parser.Parse(os.Args)

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
