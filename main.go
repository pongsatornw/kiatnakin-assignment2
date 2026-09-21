package main

import "math"

// Merge: merge 3 slices and return output as a single slices sorted by ascending.
//
// - Assume that first slice always be sorted by descending, and others is sorted by ascending
func Merge(slice1, slice2, slice3 []int) []int {
	out := make([]int, 0, len(slice1)+len(slice2)+len(slice3))

	// force first slice to be in the same direction as other slices
	slice1Asc := make([]int, 0, len(slice1))
	for i := range slice1 {
		slice1Asc = append(slice1Asc, slice1[len(slice1)-(i+1)])
	}

	var i, j, k int

	for i < len(slice1Asc) || j < len(slice2) || k < len(slice3) {
		x := getValueFromSlice(i, slice1Asc)
		y := getValueFromSlice(j, slice2)
		z := getValueFromSlice(k, slice3)

		if x <= y && x <= z {
			out = append(out, x)
			i += 1

			continue
		}

		if y <= x && y <= z {
			out = append(out, y)
			j++
			continue
		}

		out = append(out, z)
		k += 1
	}

	return out
}

func getValueFromSlice(index int, slice []int) int {
	if index < len(slice) {
		return slice[index]
	}

	return math.MaxInt
}
