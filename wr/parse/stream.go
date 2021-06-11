package parse

type Instruct struct {
	Name string
	Args []*ParserToken
	Code uint16
}

type Refer struct {
	SubType string
	Id      string
}

type State struct {
	Name string
}

type Route struct {
	Name  string
	Body  []*ParserToken
	Index int
}

type HexLit struct {
	Value uint16
}

type CharsLit struct {
	value []byte
}

type Lit struct {
	Type  string
	Hex   HexLit
	Chars CharsLit
}

type ParserToken struct {
	Type        string
	Line        int
	MemIndex    int
	Instruction *Instruct
	Reference   *Refer
	Statement   *State
	Literal     *Lit
	Routine     *Route
}
