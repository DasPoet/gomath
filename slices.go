package gomath

import (
	"sort"
)

// UnitSlice returns a slice of length "length" containing "element" "length" times
func UnitSlice(element float64, length int) []float64 {
	unitSlice := make([]float64, length)
	for i := 0; i < length; i++ {
		unitSlice[i] = element
	}
	return unitSlice
}

// Contains returns whether a given slice of integers contains a target integer
func Contains(slice []int, target int) bool {
	sort.Ints(slice)
	ind := sort.SearchInts(slice, target)
	return ind < len(slice) && slice[ind] == target
}
