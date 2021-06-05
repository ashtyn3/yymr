package main

import (
	"fmt"
	"yymr/cpu"
	opcode "yymr/opcodes"
)

func debugRegisters(c cpu.CPU) {
	rCount := 1
	for i, r := range c.Registers {
		if i == cpu.Ip {
			fmt.Printf("IP:  0x%04x\n", r)
		}
		if i == cpu.ACC {
			fmt.Printf("ACC: 0x%04x\n", r)
		}
		if i == cpu.Sp {
			fmt.Printf("SP:  0x%04x\n", r)
		}
		if i == cpu.Fp {
			fmt.Printf("Fp:  0x%04x\n", r)
		}
		if i >= cpu.R1 {
			fmt.Printf("R%d:  0x%04x\n", rCount, r)
			rCount++
		}
	}
}

func main() {
	c := cpu.Cpu()
	ram := cpu.RamDevice(256 * 256)
	c.Memory.Map(ram)
	i := -1

	i++
	ram.Mem[i] = opcode.MovLitReg
	i++
	ram.Mem[i] = cpu.R1
	i++
	ram.Mem[i] = 0x0002

	i++
	ram.Mem[i] = opcode.MovRegReg
	i++
	ram.Mem[i] = cpu.R1
	i++
	ram.Mem[i] = cpu.R2

	i++
	ram.Mem[i] = opcode.AddRegReg
	i++
	ram.Mem[i] = cpu.R1
	i++
	ram.Mem[i] = cpu.R2

	i++
	ram.Mem[i] = opcode.Hlt

	c.Run()
	debugRegisters(*c)

}
