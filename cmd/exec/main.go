package main

import (
	"fmt"
	"os"
	"yymr/cpu"
	"yymr/drivers"
	opcode "yymr/opcodes"

	"github.com/akamensky/argparse"
)

func main() {
	p := argparse.NewParser("yymr-exec", "VM execution")

	name := p.String("i", "input", &argparse.Options{Required: true})
	debug := p.Flag("d", "debug", &argparse.Options{})

	p.Parse(os.Args)
	c := cpu.Cpu()
	ram := cpu.RamDevice(256 * 256)
	screen := drivers.ScreenDevice()
	c.Memory.Map(ram)
	c.Memory.Map(screen)

	if *name == "" {
		fmt.Println(p.Usage("Needs filename"))
		os.Exit(0)
	}

	f, _ := os.ReadFile(*name)
	mem := []uint64{}
	for _, b := range f {
		if b != 0 {
			mem = append(mem, uint64(b))
		}
	}

	for i, b := range mem {
		ram.Mem[i] = b
	}
	ram.Mem[len(mem)] = opcode.Opcodes["Hlt"].Code
	c.Run()

	if *debug == true {
		c.DebugRegisters()
	}
}
