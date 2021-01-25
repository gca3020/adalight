package main

import (
	"github.com/gca3020/adalight"
	"github.com/gca3020/adalight/effects"
	"image/color"
	"time"
)

func main() {
	c := adalight.New("/dev/ttyUSB0", 115200, 106)
	c.Run(effects.NewFlash(color.NRGBA{B: 0x7F, A: 0xFF}, 2*time.Second))
	c.Run(effects.NewTrain(color.NRGBA{G: 0xFF, A: 0xFF}, 2))
	c.Run(effects.NewRamp(color.NRGBA{R: 0xFF}, 1*time.Second))
}
