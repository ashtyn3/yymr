package opcode

const (
	MovLitReg = 0x10
	MovLitMem = 0x11

	MovRegPtrReg = 0x12
	MovLitAReg   = 0x13

	MovRegReg = 0x14
	MovRegMem = 0x15
	MovMemReg = 0x16

	// Math
	AddRegReg = 0x40
	MulRegReg = 0x41

	DivRegReg = 0x42

	IncReg = 0x43
	DecReg = 0x44

	// conditionals
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

	// stack
	PshLit = 0x31
	PshReg = 0x32
	Pop    = 0x33

	// sub-routine
	CalLit = 0x34
	CalReg = 0x35
	Ret    = 0x36

	// sys-calls
	Hlt = 0x50
)
