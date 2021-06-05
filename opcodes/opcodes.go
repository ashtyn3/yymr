package opcode

const (
	MovLitReg    = 0x10
	MovRegReg    = 0x11
	MovRegMem    = 0x12
	MovMemReg    = 0x13
	AddRegReg    = 0x14
	JmpNotEq     = 0x21
	JmpEq        = 0x22
	JmpLessEq    = 0x23
	JmpLess      = 0x24
	JmpGreaterEq = 0x25
	JmpGreater   = 0x26
	PshLit       = 0x31
	PshReg       = 0x32
	Pop          = 0x33
	CalLit       = 0x34
	CalReg       = 0x35
	Ret          = 0x36
	Hlt          = 0x40
)
