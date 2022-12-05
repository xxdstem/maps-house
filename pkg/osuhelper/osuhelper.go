package osuhelper

import (
	"fmt"
	"math"
	"strconv"
)

func ModeToStr(mode int8) string {
	if mode == 3 {
		return "mania"
	}
	if mode == 2 {
		return "ctb"
	}
	if mode == 1 {
		return "taiko"
	}
	return "std"
}

func LevelProgress(l float64) float64 {
	_, f := math.Modf(l)
	f *= 100
	p, _ := strconv.ParseFloat(fmt.Sprintf("%.0f", f), 64)
	return p
}
