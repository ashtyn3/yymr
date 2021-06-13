package memory

import "log"

type Device interface {
	Start() uint
	End() uint
	GetInt(addr uint32) uint32
	SetInt(addr uint32, value uint32)
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

func (M *Mapper) GetInt(addr uint32) uint32 {
	device := M.find(uint(addr))
	if device == nil {
		log.Fatal("Could not find mapping for address: ", addr)
	}
	if device.Remap() == true {
		addr = addr - uint32(device.Start())
	}
	return device.GetInt(addr)
}

func (M *Mapper) SetInt(addr uint32, value uint32) {
	device := M.find(uint(addr))
	if device.Remap() == true {
		addr = addr - uint32(device.Start())
	}
	device.SetInt(addr, value)
}
