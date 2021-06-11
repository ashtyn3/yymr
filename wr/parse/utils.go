package parse

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"yymr/wr/token"
)

var keywords = map[string]string{
	"movLR": "MovLitReg",
	"movLM": "MovLitMem",

	"movRPR": "MovRegPtrReg",
	"movLAR": "MovLitAReg",

	"movRR": "MovRegReg",
	"movRM": "MovRegMem",
	"movMR": "MovMemReg",

	// Math
	"addRR": "AddRegReg",
	"mulRR": "MulRegReg",
	"divRR": "DivRegReg",

	"incR": "IncReg",
	"decR": "DecReg",

	// conditionals
	"jmpNE":  "JmpNotEq",
	"jmpNER": "JmpNotRegEq",

	"jmpE":  "JmpEq",
	"jmpER": "JmpRegEq",

	"jmpLE":  "JmpLessEq",
	"jmpLER": "JmpLessRegEq",

	"jmpL":  "JmpLess",
	"jmpLR": "JmpLessReg",

	"jmpGE":  "JmpGreaterEq",
	"jmpGER": "JmpGreaterRegEq",

	"jmpG":  "JmpGreater",
	"jmpGR": "JmpRegGreater",

	// stack
	"pushL": "PshLit",
	"pushR": "PshReg",
	"pop":   "Pop",

	// sub-routine
	"callL": "CalLit",
	"callR": "CalReg",
	"ret":   "Ret",

	// sys-calls
	"hlt": "Hlt",
}

func ArgToId(T int) (str string) {
	if T == token.Hex {
		str += "L"
	} else if T == token.String {
		str += "L"
	} else if T == token.MemId {
		str += "M"
	} else if T == token.RegisterId {
		str += "R"
	}

	return str
}

func KeywordToCode(s string, args []token.Token) string {
	str := s
	if strings.HasPrefix(s, "mov") || strings.HasPrefix(s, "jmp") {
		if len(args) < 2 {
			fmt.Println(strconv.Itoa(Tok.Line) + ": too few arguements for " + s + " instruction")
			os.Exit(0)
		}
		for _, t := range args {
			str += ArgToId(t.Type)
		}
	}
	if strings.HasPrefix(s, "call") || strings.HasPrefix(s, "inc") || strings.HasPrefix(s, "dec") {
		if len(args) < 1 {
			fmt.Println(strconv.Itoa(Tok.Line) + ": too few arguements for " + s + " instruction")
			os.Exit(0)
		}
		for _, t := range args {
			str += ArgToId(t.Type)
		}
	}
	return keywords[str]
}
