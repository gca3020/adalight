package effects

import (
	"github.com/gca3020/adalight"
	"image/color"
)

// Train is a simple effects that lights up every light in sequence a fixed number of times
type Train struct {
	color   color.Color
	repeats int
	strip   *adalight.Strip
}

// NewTrain builds and returns a new Train Effect
func NewTrain(c color.Color, r int) (f *Train) {
	return &Train{color: c, repeats: r}
}

// Init initializes this Effect, passing it the Strip
func (t *Train) Init(s *adalight.Strip, _ int) string {
	t.strip = s
	return "Train"
}

// Frame populates the pixels in the strip passed during initialization
func (t *Train) Frame(num int) bool {
	// Set the strip to black
	t.strip.SetAllRGB(color.RGBA{})

	// Light up the single pixel
	if num < t.repeats*t.strip.Length() {
		t.strip.Set(num%t.strip.Length(), t.color)
		return false
	}
	return true
}
