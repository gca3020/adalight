package adalight

import (
	"fmt"
	"image/color"
)

// Strip represents an LED strip, and contains methods for interacting with the LEDs on the strip
type Strip struct {
	length int
	pixels []uint8
}

// Length returns the length of the strip (number of LEDs)
func (s *Strip) Length() int {
	return s.length
}

// SetRGB sets a specific pixel `i` to a given color.NRGBA color value
func (s *Strip) SetRGB(i int, rgb color.NRGBA) {
	if i < 0 || i > s.length {
		return
	}
	s.pixels[(3*i)+0] = rgb.R
	s.pixels[(3*i)+1] = rgb.G
	s.pixels[(3*i)+2] = rgb.B
}

// Set sets a specific pixel `i` to a given color value
func (s *Strip) Set(i int, c color.Color) {
	rgb := color.NRGBAModel.Convert(c).(color.NRGBA)
	s.SetRGB(i, rgb)
}

// SetAllRGB sets all of the pixels on the strip to the given color.NRGBA value
func (s *Strip) SetAllRGB(rgb color.NRGBA) {
	for i := 0; i < s.length; i++ {
		s.SetRGB(i, rgb)
	}
}

// SetAll sets all of the pixels on the strip to the given color.Color value
func (s *Strip) SetAll(c color.Color) {
	rgb := color.NRGBAModel.Convert(c).(color.NRGBA)
	s.SetAllRGB(rgb)
}

func newStrip(length int) *Strip {
	if length < 1 {
		fmt.Println("Invalid Strip length:", length)
		panic(fmt.Sprintf("%v", length))
	}
	s := &Strip{length: length, pixels: make([]uint8, length*3)}
	return s
}
