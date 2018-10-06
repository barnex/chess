package chess

type Color int8

const (
	White  Color = 1
	Black  Color = -1
	Nobody Color = 0
)

func (c Color) String() string {
	switch c {
	case White:
		return "white"
	case Black:
		return "black"
	default:
		return "nobody"
	}
}
