package types

type SmallNumber struct {
	One   int8
	Two   int
	Three float32
	Four  float64
}

type SmallString struct {
	One   string
	Two   string
	Three string
	Four  string
}

type SmallMixed struct {
	One   int
	Two   string
	Three bool
	Four  []string
}
