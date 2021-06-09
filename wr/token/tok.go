package token

const (
	Id         = 1
	Hex        = 2
	Int        = 3
	String     = 4
	Keyword    = 5
	LeftParen  = 6
	RightParen = 7
	LeftBrack  = 8
	RightBrack = 9
	RegisterId = 10
	MemId      = 11
	Unknown    = 12
)

type Token struct {
	Type int
	Line int
	Col  int
	Text string
}
