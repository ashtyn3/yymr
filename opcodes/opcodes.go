package opcode

var Opcodes = map[string]Opcode{
	"MovLitReg": {Code: 0x10, Size: 3},
	"MovLitMem": {Code: 0x11, Size: 3},

	"MovRegPtrReg": {Code: 0x12, Size: 3},
	"MovLitAReg":   {Code: 0x13, Size: 4},

	"MovRegReg": {Code: 0x14, Size: 3},
	"MovRegMem": {Code: 0x15, Size: 3},
	"MovMemReg": {Code: 0x16, Size: 3},

	// Math
	"AddRegReg": {Code: 0x40, Size: 3},
	"MulRegReg": {Code: 0x41, Size: 3},
	"DivRegReg": {Code: 0x42, Size: 3},

	"IncReg": {Code: 0x43, Size: 2},
	"DecReg": {Code: 0x44, Size: 2},

	// conditionals
	"JmpNotEq":    {Code: 0x21, Size: 3},
	"JmpNotRegEq": {Code: 0x22, Size: 3},

	"JmpEq":    {Code: 0x23, Size: 3},
	"JmpRegEq": {Code: 0x24, Size: 3},

	"JmpLessEq":    {Code: 0x25, Size: 3},
	"JmpLessRegEq": {Code: 0x26, Size: 3},

	"JmpLess":    {Code: 0x27, Size: 3},
	"JmpLessReg": {Code: 0x28, Size: 3},

	"JmpGreaterEq":    {Code: 0x29, Size: 3},
	"JmpGreaterRegEq": {Code: 0x2a, Size: 3},

	"JmpGreater":    {Code: 0x2b, Size: 3},
	"JmpRegGreater": {Code: 0x2c, Size: 3},

	// stack
	"PshLit": {Code: 0x31, Size: 2},
	"PshReg": {Code: 0x32, Size: 2},
	"Pop":    {Code: 0x33, Size: 2},

	// sub-routine
	"CalLit": {Code: 0x34, Size: 2},
	"CalReg": {Code: 0x35, Size: 2},
	"Ret   ": {Code: 0x36},

	// sys-calls
	"Hlt": {Code: 0x50, Size: 1},
}
