package day6

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Lab struct {
	width, height int
	obstructions  mapset.Set[Point]
}
