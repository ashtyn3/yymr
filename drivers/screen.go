package drivers

import (
	"fmt"
	"math"

	"github.com/ahmetb/go-cursor"
)

type Screen struct{}

func (r *Screen) Start() uint {
	return 0x3000
}
func (r *Screen) End() uint {
	return 0x30ff
}
func (r *Screen) Remap() bool {
	return true
}

func (r *Screen) GetInt(addr uint32) uint32 {
	return 0x0000
}

func (r *Screen) SetInt(addr uint32, value uint32) {
	//command := (value & 0xff00) >> 8
	characterValue := value & 0x00ff

	x := (addr % 16) + 1
	y := math.Floor(float64(addr/16)) + 1
	character := string(rune(int(characterValue)))
	fmt.Print(cursor.MoveTo(int(y), int(x)), character)
}

func ScreenDevice() *Screen {
	r := Screen{}

	return &r
}
