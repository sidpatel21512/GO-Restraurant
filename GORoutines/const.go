package main

import (
	"sync"
)

var availableTables []int
var bookedTables []int

var customerCount = 0
var maxCustomers = 5
var customerQueue []Customer

var minSec = 1
var maxSec = 5

var minDiningTime = 10
var maxDiningTime = 15

var restaurantTiming = 20
var isRestaurantTimeOver = false
var isRestaurantClosed = false

var mut = &sync.Mutex{}

func init() {
	availableTables = []int{1, 2}
	bookedTables = []int{}
	customerQueue = []Customer{}
}
