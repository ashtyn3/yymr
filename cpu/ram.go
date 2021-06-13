package cpu

type Ram struct {
	Mem []uint32
}

func (r *Ram) Remap() bool {
	return false
}
func (r *Ram) Start() uint {
	return 0x0000
}
func (r *Ram) End() uint {
	return 0xffff
}

func (r *Ram) GetInt(addr uint32) uint32 {
	return r.Mem[addr]
}

func (r *Ram) SetInt(addr uint32, value uint32) {
	r.Mem[addr] = value
}

func RamDevice(size int) *Ram {
	r := Ram{}

	r.Mem = make([]uint32, size)

	return &r
}
