package colour

type Colour interface {
	To8Bit() RGB8Bit
}

type RGB8Bit struct {
	R, G, B byte // We are using 8-bit colour
}

func New8Bit(r, g, b byte) RGB8Bit {
	return RGB8Bit{r, g, b}
}

func (eightBit RGB8Bit) To8Bit() RGB8Bit {
	return eightBit
}

type RGB struct {
	R, G, B float64
}

func NewRGB(r, g, b float64) RGB {
	if !onInterval(0, 1, r) || !onInterval(0, 1, g) || !onInterval(0, 1, b) {
		panic("RGB intensity values must be on the interval [0, 1].")
	}

	return RGB{
		R: r,
		G: g,
		B: b,
	}
}

func onInterval(lower, upper, value float64) bool {
	return lower <= value && value <= upper
}

func (rgb RGB) To8Bit() RGB8Bit {
	return RGB8Bit{
		R: byte(rgb.R * 255),
		G: byte(rgb.G * 255),
		B: byte(rgb.B * 255),
	}
}
