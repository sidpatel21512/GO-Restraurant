package main

import (
	"math/rand"
	"strconv"
)

func generateRandomNumber(min int, max int) int {
	return rand.Intn(max) + min
}

func indexOf(data []int, ele int) int {
	for k, v := range data {
		if ele == v {
			return k
		}
	}
	return -1 //not found.
}

func parseInt(value string) int {
	parseValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return parseValue
}
