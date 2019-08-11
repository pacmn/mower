package mow

type Orientation uint32

const (
	N Orientation = 0
	E Orientation = 1
	S Orientation = 2
	W Orientation = 3
)

type Mower struct {
	X           int
	Y           int
	Orientation Orientation
	Movements   string
}
