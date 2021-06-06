package memory

import "log"

type Device interface {
	Start() uint
	End() uint
	GetInt(addr uint16) uint16
	SetInt(addr uint16, value uint16)
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

func (M *Mapper) GetInt(addr uint16) uint16 {
	device := M.find(uint(addr))
	if device == nil {
		log.Fatal("Could not find mapping for address: ", addr)
	}
	if device.Remap() == true {
		addr = addr - uint16(device.Start())
	}
	return device.GetInt(addr)
}

func (M *Mapper) SetInt(addr uint16, value uint16) {
	device := M.find(uint(addr))
	if device.Remap() == true {
		addr = addr - uint16(device.Start())
	}
	device.SetInt(addr, value)
}
