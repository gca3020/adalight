package effects

import (
	"github.com/gca3020/adalight"
	"image/color"
	"time"
)

// Flash is a simple effects that ramps up and down a single color over a given duration
type Flash struct {
	color  color.Color
	dur    time.Duration
	strip  *adalight.Strip
	frames int
}

// NewFlash builds and returns a new Flash effect
func NewFlash(c color.Color, d time.Duration) (f *Flash) {
	return &Flash{color: c, dur: d}
}

// Init initializes this Effect, passing it the Strip
func (f *Flash) Init(s *adalight.Strip, fps int) string {
	f.strip = s
	f.frames = int(float64(fps) * f.dur.Seconds())
	return "Flash"
}

// Frame populates the pixels in the strip passed during initialization
func (f *Flash) Frame(num int) bool {
	if num < f.frames {
		f.strip.SetAll(f.color)
		return false
	} else {
		f.strip.SetAllRGB(color.RGBA{})
		return true
	}
}
