package main

import "math/rand"

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}

func indexOf(data []int, ele int) int {
	for k, v := range data {
		if ele == v {
			return k
		}
	}
	return -1 //not found.
}
