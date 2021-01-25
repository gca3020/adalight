package effects

import (
	"github.com/gca3020/adalight"
	"image/color"
	"time"
)

// Ramp is a simple effects that ramps up and down a single color over a given duration
type Ramp struct {
	color  color.Color
	dur    time.Duration
	strip  *adalight.Strip
	frames int
}

// NewRamp builds and returns a new Ramp effect
func NewRamp(c color.Color, d time.Duration) (r *Ramp) {
	return &Ramp{color: c, dur: d}
}

// Init initializes this Effect, passing it the Strip
func (r *Ramp) Init(s *adalight.Strip, fps int) string {
	r.strip = s
	r.frames = int(float64(fps)*r.dur.Seconds()) / 2
	return "Ramp"
}

// Frame populates the pixels in the strip passed during initialization
func (r *Ramp) Frame(num int) bool {
	// Early return if we're done
	if num >= 2*r.frames {
		r.strip.SetAllRGB(color.RGBA{})
		return true
	}

	// Calculate an alpha component based on the frame number
	alpha := uint8((float64(num%r.frames) / float64(r.frames)) * 255.0)
	if num >= r.frames && num < 2*r.frames {
		alpha = 255 - alpha
	}

	// Generate a new color by scaling the given color by the alpha component
	rgb := color.NRGBAModel.Convert(r.color).(color.NRGBA)
	rgb.A = alpha

	r.strip.SetAll(rgb)
	return false
}
