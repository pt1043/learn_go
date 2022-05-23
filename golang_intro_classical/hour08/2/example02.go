package main

import (
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

func (s *Sphere) surfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

func main() {
	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.surfaceArea())
	fmt.Println(s.Volume())
}
