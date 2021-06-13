package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"yymr/cpu"
	opcode "yymr/opcodes"

	"github.com/akamensky/argparse"
)

func printChar(char int, ram *cpu.Ram, i *int, pos int) *cpu.Ram {
	*i++
	ram.Mem[*i] = opcode.Opcodes["MovLitReg"].Code
	*i++
	ram.Mem[*i] = cpu.R1
	*i++
	ram.Mem[*i] = uint64(char)

	*i++
	ram.Mem[*i] = opcode.Opcodes["MovRegMem"].Code
	*i++
	ram.Mem[*i] = cpu.R1
	*i++
	ram.Mem[*i] = uint64(0x3000 + pos)
	return ram
}

func tempId() string {
	b, _ := rand.Prime(rand.Reader, 100)
	s := sha256.New()
	s.Write(b.Bytes())

	return hex.EncodeToString(s.Sum(nil))[0:15]
}

func main() {
	p := argparse.NewParser("yymr", "A fast virtual machine")

	f := p.File("I", "input", os.O_RDWR, 0755, &argparse.Options{Required: true, Help: "Specifies input file"})
	v := p.Flag("v", "verbose", &argparse.Options{Help: "Shows build steps"})
	flags := p.String("F", "flags", &argparse.Options{})

	p.Parse(os.Args)

	if *f == *new(os.File) {
		fmt.Print(p.Usage("Needs filename"))
		os.Exit(0)
	}

	tmp := os.TempDir()
	id := tempId()
	name, _ := filepath.Abs(f.Name())
	os.Setenv("WORK", tmp+"/"+id)

	os.Mkdir(os.Getenv("WORK"), 0755)
	os.Chdir(os.Getenv("WORK"))

	if *v {
		fmt.Println("Work Dir: " + os.Getenv("WORK"))
		fmt.Println("Changed Dir:" + os.Getenv("WORK"))
	}

	c := exec.Command("yymr-pt", "-i", name, "-o", os.Getenv("WORK")+"/a.a")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr

	cErr := c.Run()

	if cErr != nil {
		fmt.Println(cErr)
	}

	if *v {
		fmt.Println("Pre-processing Source: " + name)
		fmt.Println(c.String())
	}
	flagArr := []string{"-i", os.Getenv("WORK") + "/a.a"}
	if *flags != "" {
		flagArr = append(flagArr, strings.Split(strings.TrimSpace(*flags), ",")...)
	}
	e := exec.Command("yymr-exec", flagArr...)
	e.Stdout = os.Stdout
	e.Stdin = os.Stdin
	e.Stderr = os.Stderr

	if *v {
		fmt.Println("Running: " + os.Getenv("WORK") + "/a.a")
		fmt.Println(e.String())
	}

	err := e.Run()

	if err != nil {
		fmt.Println(cErr)
	}
}
