package cpu

type Ram struct {
	Mem []uint16
}

func (r *Ram) Start() uint {
	return 0x0000
}
func (r *Ram) End() uint {
	return 0xffff
}

func (r *Ram) GetInt(addr uint16) uint16 {
	return r.Mem[addr]
}

func (r *Ram) SetInt(addr uint16, value uint16) {
	r.Mem[addr] = value
}

func RamDevice(size int) *Ram {
	r := Ram{}

	for i := 0; i < size; i++ {
		r.Mem = append(r.Mem, 0x0000)
	}

	return &r
}
