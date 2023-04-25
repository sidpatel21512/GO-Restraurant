package main

import (
	"fmt"
	"time"
)

func startCustomerGeneration() {
	for custCount < maxOrders {
		c := generateCustomer()
		fmt.Printf("Customer %v has arrived at the restaurant. \n", c)

		assignTable(c)
	}

	fmt.Println("We are now closing and not taking more customers !!!")

	shutDown()
}

func generateCustomer() int {
	t := generateRandomNumber(minSec, maxSec)
	time.Sleep(time.Duration(t) * time.Second)
	custCount++
	return custCount
}

func addCustomerInQueue(c int) {
	cQueue = append(cQueue, c)
}

func removeCustomerFromQueue() int {
	cust := cQueue[0]
	cQueue = cQueue[1:]
	return cust
}

func shutDown() {
	cLen := len(cQueue)
	time.Sleep(time.Duration(cLen*maxDiningPeriod) * time.Second)
	fmt.Println("We are now closed.")
}
