package colour

type Colour struct {
	R, G, B byte // We are using 8-bit colour
}

func New(r, g, b byte) Colour {
	return Colour{r, g, b}
}
