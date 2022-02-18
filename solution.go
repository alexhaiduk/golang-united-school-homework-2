package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type Sides int

const (
	SidesCircle   Sides = 0
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
)

//func CalcSquare(sideLen float64, sidesNum #yourTypeNameHere#) float64 {
func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	//var result float64
	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == SidesTriangle {
		return (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	} else {
		return 0
	}
	/*
		switch sidesNum {
		default:
			return 0
		case SidesCircle:
			return math.Pi * math.Pow(sideLen, 2)
		case SidesTriangle:
			return (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
		case SidesSquare:
			return math.Pow(sideLen, 2)
		}
	*/
}
