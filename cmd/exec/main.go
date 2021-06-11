package main

import (
	"fmt"
	"os"
	"yymr/cpu"
	"yymr/drivers"
	opcode "yymr/opcodes"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("No file specified")
		os.Exit(0)
	}
	c := cpu.Cpu()
	ram := cpu.RamDevice(256 * 256)
	screen := drivers.ScreenDevice()
	c.Memory.Map(ram)
	c.Memory.Map(screen)

	f, _ := os.ReadFile(os.Args[1])
	mem := []uint16{}
	for _, b := range f {
		if b != 0 {
			mem = append(mem, uint16(b))
		}
	}

	for i, b := range mem {
		ram.Mem[i] = b
	}
	ram.Mem[len(mem)] = opcode.Opcodes["Hlt"].Code
	c.Run()
	c.DebugRegisters()
}
