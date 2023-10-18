package main

import (
	"math/rand"
)

func Merge(arr []int, left []int, right []int) {
	var i, j, k int
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			arr[k] = right[j]
			j += 1
		} else if left[i] < right[j] {

			arr[k] = left[i]
			i += 1
		} else {
			arr[k] = left[i]
			k += 1
			arr[k] = left[i]
			i += 1
			j += 1

		}
		k += 1
	}

	for i < len(left) {
		arr[k] = left[i]
		i += 1
		k += 1

	}

	for j < len(right) {
		arr[k] = right[j]
		j += 1
		k += 1

	}
}


func GenerateRandomSlice(length int) []int {
	res := make([]int, length)
	i := 0
	for i < length {
		number := rand.Intn(500000)
		res = append(res, number)
		i += 1
	}
	return res
}

