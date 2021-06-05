package opcode

const (
	MovLitReg = 0x10
	MovRegReg = 0x11
	MovRegMem = 0x12
	MovMemReg = 0x13
	AddRegReg = 0x14

	JmpNotEq    = 0x21
	JmpNotRegEq = 0x22

	JmpEq    = 0x23
	JmpRegEq = 0x24

	JmpLessEq    = 0x25
	JmpLessRegEq = 0x26

	JmpLess    = 0x27
	JmpLessReg = 0x28

	JmpGreaterEq    = 0x29
	JmpGreaterRegEq = 0x2a

	JmpGreater    = 0x2b
	JmpRegGreater = 0x2c

	PshLit = 0x31
	PshReg = 0x32
	Pop    = 0x33
	CalLit = 0x34
	CalReg = 0x35
	Ret    = 0x36
	Hlt    = 0x40
)
