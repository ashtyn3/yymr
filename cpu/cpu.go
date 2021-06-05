package cpu

import (
	"fmt"
	"log"
	"yymr/cpu/memory"
	opcode "yymr/opcodes"
)

type CPU struct {
	Registers [12]uint16
	Memory    *memory.Mapper
}

func (c *CPU) getRegister(r uint8) uint16 {
	if r > 12 {
		log.Fatalln("Failed to get register:", r)
	}
	return c.Registers[r]
}

func (c *CPU) setRegister(r uint8, value uint16) {
	if r > 12 {
		log.Fatalln("Failed to get register:", r)
	}
	c.Registers[r] = value
}

func Cpu() *CPU {
	c := &CPU{}
	c.Memory = &memory.Mapper{}

	return c
}

func (c *CPU) Fetch() uint16 {
	pointer := c.getRegister(Ip)
	ip := c.Memory.GetInt(pointer)
	c.setRegister(Ip, pointer+1)

	return ip
}

func (c *CPU) Execute(op uint16) {
	fmt.Println(op)
	switch op {
	case opcode.MovLitReg:
		{
			reg := c.Fetch()
			val := c.Fetch()

			c.setRegister(uint8(reg), val)
			break
		}
	case opcode.MovRegReg:
		{
			from := c.Fetch()
			val := c.getRegister(uint8(from))
			to := c.Fetch()

			c.setRegister(uint8(to), val)
			break
		}
	case opcode.MovMemReg:
		{
			fromMem := c.Fetch()
			to := c.Fetch()
			val := c.Memory.GetInt(fromMem)

			c.setRegister(uint8(to), val)
			break
		}
	case opcode.MovRegMem:
		{
			fromReg := c.Fetch()
			toMem := c.Fetch()
			val := c.getRegister(uint8(fromReg))

			c.Memory.SetInt(toMem, val)
			break
		}
	case opcode.AddRegReg:
		{
			r1 := c.Fetch()
			r2 := c.Fetch()
			val1 := c.getRegister(uint8(r1))
			val2 := c.getRegister(uint8(r2))

			c.setRegister(ACC, val1+val2)
			break
		}
	}
}

func (c *CPU) Step() {
	op := c.Fetch()
	c.Execute(op)
}
