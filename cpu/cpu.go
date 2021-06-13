package cpu

import (
	"fmt"
	"log"
	"yymr/cpu/memory"
	opcode "yymr/opcodes"
)

type CPU struct {
	Registers [12]uint32
	Memory    *memory.Mapper
	frameSize uint32
}

func (c *CPU) getRegister(r uint32) uint32 {
	if int(r) > len(c.Registers) {
		log.Fatalln("Failed to get register:", r)
	}
	return c.Registers[r]
}

func (c *CPU) setRegister(r uint32, value uint32) {
	if int(r) > len(c.Registers) {
		log.Fatalln("Failed to get register:", r)
	}
	c.Registers[r] = value
}

func Cpu() *CPU {
	c := &CPU{}
	c.Memory = &memory.Mapper{}
	c.Registers[2] = uint32(len(c.Memory.Ranges) - 1)
	c.Registers[3] = uint32(len(c.Memory.Ranges) - 1)

	return c
}

func (c *CPU) Fetch() uint32 {
	pointer := c.getRegister(Ip)
	ip := c.Memory.GetInt(pointer)
	c.setRegister(Ip, pointer+1)

	return ip
}

func (c *CPU) Execute(op uint32) int {
	switch op {
	// Move instructions
	case opcode.Opcodes["MovLitReg"].Code:
		{
			val := c.Fetch()
			reg := c.Fetch()

			c.setRegister(uint32(reg), val)
			break
		}
	case opcode.Opcodes["MovRegReg"].Code:
		{
			from := c.Fetch()
			val := c.getRegister(uint32(from))
			to := c.Fetch()

			c.setRegister(uint32(to), val)
			break
		}
	case opcode.Opcodes["MovMemReg"].Code:
		{
			fromMem := c.Fetch()
			to := c.Fetch()
			val := c.Memory.GetInt(fromMem)

			c.setRegister(uint32(to), val)
			break
		}
	case opcode.Opcodes["MovRegMem"].Code:
		{
			fromReg := c.Fetch()
			toMem := c.Fetch()
			val := c.getRegister(uint32(fromReg))

			c.Memory.SetInt(toMem, val)
			break
		}
	case opcode.Opcodes["MovLitMem"].Code:
		{
			lit := c.Fetch()
			toMem := c.Fetch()

			c.Memory.SetInt(toMem, lit)
			break
		}
	case opcode.Opcodes["MovRegPtrReg"].Code:
		{
			rFrom := c.Fetch()
			rTo := c.Fetch()
			addr := c.getRegister(uint32(rFrom))
			val := c.Memory.GetInt(addr)

			c.setRegister(uint32(rTo), val)
			break
		}
	case opcode.Opcodes["MovLitAReg"].Code:
		{
			base := c.Fetch()
			offset := c.getRegister(uint32(c.Fetch()))
			to := c.Fetch()

			val := c.Memory.GetInt(base + offset)

			c.setRegister(uint32(to), val)

			break
		}
		// Math instructions
	case opcode.Opcodes["AddRegReg"].Code:
		{
			r1 := c.Fetch()
			r2 := c.Fetch()
			val1 := c.getRegister(uint32(r1))
			val2 := c.getRegister(uint32(r2))

			c.setRegister(ACC, val1+val2)
			break
		}
	case opcode.Opcodes["MulRegReg"].Code:
		{
			r1 := c.Fetch()
			r2 := c.Fetch()
			val1 := c.getRegister(uint32(r1))
			val2 := c.getRegister(uint32(r2))

			c.setRegister(ACC, val1*val2)
			break
		}
	case opcode.Opcodes["opcode.IncReg"].Code:
		{
			r := c.Fetch()
			val := c.getRegister(uint32(r)) + 1

			c.setRegister(uint32(r), val)
			break
		}
	case opcode.Opcodes["DecReg"].Code:
		{
			r := c.Fetch()
			val := c.getRegister(uint32(r)) - 1

			c.setRegister(uint32(r), val)
			break
		}
	case opcode.Opcodes["opcode.DivRegReg"].Code:
		{
			r1 := c.Fetch()
			r2 := c.Fetch()
			val1 := c.getRegister(uint32(r1))
			val2 := c.getRegister(uint32(r2))

			c.setRegister(ACC, val1/val2)
			break
		}
	// Jump instructions
	case opcode.Opcodes["opcode.JmpEq"].Code:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit == c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpRegEq"].Code:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint32(reg)) == c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpNotEq"].Code:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit != c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpNotRegEq"].Code:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint32(reg)) != c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpLessEq"].Code:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit <= c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpLessRegEq"].Code:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint32(reg)) <= c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["OpcodesJmpLess"].Code:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit < c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpLessReg"].Code:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint32(reg)) < c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpGreaterEq"].Code:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit >= c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpGreaterRegEq"].Code:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint32(reg)) >= c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpGreater"].Code:
		{
			lit := c.Fetch()
			pointer := c.Fetch()

			if lit > c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}
	case opcode.Opcodes["JmpRegGreater"].Code:
		{
			reg := c.Fetch()
			pointer := c.Fetch()

			if c.getRegister(uint32(reg)) > c.getRegister(ACC) {
				c.setRegister(Ip, pointer)
			}
			break
		}

	// Stack
	case opcode.Opcodes["PshLit"].Code:
		{
			val := c.Fetch()
			sp := c.getRegister(Sp)

			c.Memory.SetInt(sp, val)
			c.setRegister(Sp, sp-1)
			break
		}
	case opcode.Opcodes["PshReg"].Code:
		{
			val := c.getRegister(uint32(c.Fetch()))
			c.push(val)
			break
		}
	case opcode.Opcodes["Pop"].Code:
		{
			r := c.Fetch()
			val := c.pop()
			c.setRegister(uint32(r), val)
			break
		}
		// sub-routines

	case opcode.Opcodes["CalLit"].Code:
		{
			addr := c.Fetch()
			c.saveState()

			c.setRegister(Ip, addr)
		}
	case opcode.Opcodes["CalReg"].Code:
		{
			addr := c.getRegister(uint32(c.Fetch()))
			c.saveState()

			c.setRegister(Ip, addr)
		}
	case opcode.Opcodes["Ret"].Code:
		{
			addr := c.getRegister(uint32(c.Fetch()))
			c.popState()

			c.setRegister(Ip, addr)
		}
	case opcode.Opcodes["Hlt"].Code:
		return 1
	}
	return 0
}

func (c *CPU) Step() int {
	op := c.Fetch()
	return c.Execute(op)
}

func (c *CPU) DebugRegisters() {
	rCount := 1
	for i, r := range c.Registers {
		if i == Ip {
			fmt.Printf("IP:  0x%04x\n", r)
		}
		if i == ACC {
			fmt.Printf("ACC: 0x%04x\n", r)
		}
		if i == Sp {
			fmt.Printf("SP:  0x%04x\n", r)
		}
		if i == Fp {
			fmt.Printf("Fp:  0x%04x\n", r)
		}
		if i >= R1 {
			fmt.Printf("R%d:  0x%04x\n", rCount, r)
			rCount++
		}
	}
}

func (c *CPU) Run() {
	for {
		if c.Step() == 1 {
			break
		}
	}
}
