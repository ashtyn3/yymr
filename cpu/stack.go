package cpu

import "fmt"

func (c *CPU) push(val uint64) {
	sp := c.getRegister(Sp)
	fmt.Println(sp)

	c.Memory.SetInt(sp, val)
	c.setRegister(Sp, sp-1)
	c.frameSize++
}

func (c *CPU) pop() uint64 {
	next := c.getRegister(Sp) + 1
	c.setRegister(Sp, next)
	c.frameSize++
	return c.Memory.GetInt(next)
}

func (c *CPU) saveState() {
	c.push(c.getRegister(R1))
	c.push(c.getRegister(R2))
	c.push(c.getRegister(R3))
	c.push(c.getRegister(R4))
	c.push(c.getRegister(R5))
	c.push(c.getRegister(R6))
	c.push(c.getRegister(R7))
	c.push(c.getRegister(R8))
	c.push(c.getRegister(Ip))
	c.push(c.getRegister(Sp))
	c.push(c.getRegister(Fp))
	c.push(c.frameSize + 1)

	c.setRegister(Fp, c.getRegister(Sp))
	c.frameSize = 0
}

func (c *CPU) popState() {
	fpa := c.getRegister(Fp)
	c.setRegister(Sp, fpa)

	c.frameSize = c.pop()
	oldFrameSize := c.frameSize

	c.setRegister(R1, c.pop())
	c.setRegister(R2, c.pop())
	c.setRegister(R3, c.pop())
	c.setRegister(R4, c.pop())
	c.setRegister(R5, c.pop())
	c.setRegister(R6, c.pop())
	c.setRegister(R7, c.pop())
	c.setRegister(R8, c.pop())
	c.setRegister(Ip, c.pop())
	c.setRegister(Sp, c.pop())
	c.setRegister(Fp, c.pop())

	nArgs := c.pop()

	for i := 0; i < int(nArgs); i++ {
		c.pop()
	}

	c.setRegister(Fp, fpa+oldFrameSize)
}
