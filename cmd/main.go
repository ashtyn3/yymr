package main

import (
	"yymr/cpu"
	"yymr/drivers"
	opcode "yymr/opcodes"
)

func printChar(char int, ram *cpu.Ram, i *int, pos int) *cpu.Ram {
	*i++
	ram.Mem[*i] = opcode.MovLitReg
	*i++
	ram.Mem[*i] = cpu.R1
	*i++
	ram.Mem[*i] = uint16(char)

	*i++
	ram.Mem[*i] = opcode.MovRegMem
	*i++
	ram.Mem[*i] = cpu.R1
	*i++
	ram.Mem[*i] = uint16(0x3000 + pos)
	return ram
}

func main() {
	c := cpu.Cpu()
	ram := cpu.RamDevice(256 * 256)
	screen := drivers.ScreenDevice()
	c.Memory.Map(ram)
	c.Memory.Map(screen)
	i := -1
	for pos, char := range "Hello Jeff!" {
		ram = printChar(int(char), ram, &i, pos)
	}

	i++
	ram.Mem[i] = opcode.Hlt

	c.Run()
}
