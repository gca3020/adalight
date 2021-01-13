package main

import (
	"github.com/gca3020/adalight"
	"github.com/gca3020/adalight/effects"
	"image/color"
	"time"
)

func main() {
	c := adalight.New("/dev/ttyUSB0", 115200, 106)
	c.Run(effects.NewFlash(color.NRGBA{B: 0x7F}, 2*time.Second))
	c.Run(effects.NewTrain(color.NRGBA{G: 0xFF}, 3))
}
