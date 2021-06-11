package assemble

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"yymr/wr/parse"
)

type Build struct {
	Filename string
	Libs     []string
	Exec     bool
	Toks     []parse.ParserToken
}
type Memory []uint16

var AddressTable = make(map[string]int)
var Mem = Memory{}
var id = -1

func routeBuild(t parse.ParserToken) {
	for _, t := range t.Routine.Body {
		if t.Type == "instruction" {
			instructBuild(*t)
		}
	}
}
func instructBuild(t parse.ParserToken) {
	in := t.Instruction
	id = t.MemIndex

	Mem = append(Mem, t.Instruction.Code)

	for _, t := range in.Args {
		if t.Type == "lit" {
			if t.Literal.Type == "hex" {
				v := t.Literal.Hex.Value
				Mem = append(Mem, v)
			}
		}
		if t.Type == "ref" {
			if t.Reference.SubType == "reg" {
				v, _ := strconv.Atoi(t.Reference.Id)
				Mem = append(Mem, uint16(v+4))
			}
		}
	}
}

func (b *Build) Assemble() {
	for _, t := range b.Toks {
		if t.Type == "route" {
			AddressTable[t.Routine.Name] = t.Routine.Index
			routeBuild(t)
		} else {
			fmt.Println(strconv.Itoa(t.Line) + ": expected sub-routine definition")
		}
	}
	f, _ := os.Create(b.Filename)
	binary.Write(f, binary.LittleEndian, Mem)

	if b.Exec {
		os.Chmod(b.Filename, 0700)
	}
}
