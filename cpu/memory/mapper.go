package memory

import "log"

type Device interface {
	Start() uint
	End() uint
	GetInt(addr uint64) uint64
	SetInt(addr uint64, value uint64)
	Remap() bool
}

type Mapper struct {
	Ranges []Device
}

func (M *Mapper) Map(d Device) {
	M.Ranges = append([]Device{d}, M.Ranges...)
}

func (M *Mapper) find(addr uint) Device {
	for _, d := range M.Ranges {
		if addr >= d.Start() && addr <= d.End() {
			return d
		}
	}

	return nil
}

func (M *Mapper) GetInt(addr uint64) uint64 {
	device := M.find(uint(addr))
	if device == nil {
		log.Fatal("Could not find mapping for address: ", addr)
	}
	if device.Remap() == true {
		addr = addr - uint64(device.Start())
	}
	return device.GetInt(addr)
}

func (M *Mapper) SetInt(addr uint64, value uint64) {
	device := M.find(uint(addr))
	if device.Remap() == true {
		addr = addr - uint64(device.Start())
	}
	device.SetInt(addr, value)
}
