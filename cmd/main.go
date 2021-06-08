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

	i++
	ram.Mem[i] = opcode.MovLitReg
	i++
	ram.Mem[i] = cpu.R1
	i++
	ram.Mem[i] = 0x0001

	i++
	ram.Mem[i] = opcode.MovLitReg
	i++
	ram.Mem[i] = cpu.R2
	i++
	ram.Mem[i] = 0x0002

	i++
	ram.Mem[i] = opcode.PshReg
	i++
	ram.Mem[i] = cpu.R1

	i++
	ram.Mem[i] = opcode.PshReg
	i++
	ram.Mem[i] = cpu.R2

	i++
	ram.Mem[i] = opcode.Pop
	i++
	ram.Mem[i] = cpu.R1

	i++
	ram.Mem[i] = opcode.Pop
	i++
	ram.Mem[i] = cpu.R2

	i++
	ram.Mem[i] = opcode.Hlt

	c.Run()
	c.DebugRegisters()

}
