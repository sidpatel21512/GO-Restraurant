package main

import (
	"fmt"
	"time"
)

func createTables(t int) {
	// availableTables = []int{1, 2, 3, 4}
	for i := 1; i <= t; i++ {
		availableTables = append(availableTables, i)
	}
}

func assignTable(c int) {
	if tLen := len(availableTables); tLen > 0 {
		table := availableTables[0]
		if tLen > 1 {
			n := generateRandomNumber(0, tLen-1)
			table = availableTables[n]
			bookedTables = append(bookedTables, table)
			availableTables = append(availableTables[:n], availableTables[n+1:]...)
		} else {
			bookedTables = append(bookedTables, table)
			availableTables = []int{}
		}
		fmt.Printf("Customer %v has been assigned a table %v. \n", c, table)
		diningPeriod := generateRandomNumber(minDiningPeriod, maxDiningPeriod)
		time.AfterFunc(time.Duration(diningPeriod)*time.Second, func() {
			freeTheTable(table)
			if len(cQueue) > 0 {
				cust := removeCustomerFromQueue()
				assignTable(cust)
			}
		})
	} else if queueLen > len(cQueue) {
		addCustomerInQueue(c)
		printStatus()
	} else {
		fmt.Printf("Customer %v has left the restaurant, due to waiting. \n", c)
	}

}

func freeTheTable(t int) {
	i := indexOf(bookedTables, t)
	bookedTables = append(bookedTables[:i], bookedTables[i+1:]...)
	availableTables = append(availableTables, t)
	fmt.Printf("Table %v is available. \n", t)
}

func printStatus() {
	fmt.Println("*---------------------------------------*")
	fmt.Println("Booked Tables: ", bookedTables)
	fmt.Println("Free Tables: ", availableTables)
	fmt.Println("Customers in waiting: ", cQueue)
	fmt.Println("*---------------------------------------*")
}
