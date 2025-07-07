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

type MediumNumber struct {
	One   SmallNumber
	Two   SmallNumber
	Three SmallNumber
	Four  SmallNumber
}
type MediumString struct {
	One   SmallString
	Two   SmallString
	Three SmallString
	Four  SmallString
}
type MediumMixed struct {
	One   SmallMixed
	Two   SmallMixed
	Three SmallMixed
	Four  SmallMixed
}

type LargeNumber struct {
	One   MediumNumber
	Two   MediumNumber
	Three MediumNumber
	Four  MediumNumber
}
type LargeString struct {
	One   MediumString
	Two   MediumString
	Three MediumString
	Four  MediumString
}
type LargeMixed struct {
	One   MediumMixed
	Two   MediumMixed
	Three MediumMixed
	Four  MediumMixed
}
