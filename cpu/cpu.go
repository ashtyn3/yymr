package cpu

import (
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

func (c *CPU) Execute(op uint16) int {
	switch op {
	// Move instructions
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
	// Math instructions
	case opcode.AddRegReg:
		{
			r1 := c.Fetch()
			r2 := c.Fetch()
			val1 := c.getRegister(uint8(r1))
			val2 := c.getRegister(uint8(r2))

			c.setRegister(ACC, val1+val2)
			break
		}
	// Jump instructions
	case opcode.JmpEq:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit == c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpRegEq:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint8(reg)) == c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpNotEq:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit != c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpNotRegEq:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint8(reg)) != c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpLessEq:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit <= c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpLessRegEq:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint8(reg)) <= c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpLess:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit < c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpLessReg:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint8(reg)) < c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpGreaterEq:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit >= c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpGreaterRegEq:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint8(reg)) >= c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpGreater:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit > c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.JmpRegGreater:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint8(reg)) > c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Hlt:
		return 1
	}
	return 0
}

func (c *CPU) Step() int {
	op := c.Fetch()
	return c.Execute(op)
}

func (c *CPU) Run() {
	for {
		if c.Step() == 1 {
			break
		}
	}
}
