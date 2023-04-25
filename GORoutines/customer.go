package main

import (
	"fmt"
	"sync"
	"time"
)

func generateCustomers(wg *sync.WaitGroup, ch chan<- int) {
	no := getRandomNumber(minSec, maxSec)
	time.AfterFunc(time.Duration(no)*time.Second, func() {
		createCustomer()
		ch <- customerCount
		mut.Lock()
		if !isRestaurantTimeOver {
			generateCustomers(wg, ch)
		} else {
			fmt.Println("We are not longer taking new Customers!!!")
		}
		mut.Unlock()
		defer wg.Done()
	})
	wg.Add(1)
}

func createCustomer() {
	time := getRandomNumber(minDiningTime, maxDiningTime)
	customerCount++
	addCustomer(time)
	fmt.Printf("Customer %v has arrived at the restaurant. \n", customerCount)
}

func addCustomer(time int) {
	mut.Lock()
	customerQueue = append(customerQueue, Customer{id: customerCount, dinnerTime: time})
	mut.Unlock()
}

func removeCustomer() Customer {
	cust := customerQueue[0]
	mut.Lock()
	customerQueue = customerQueue[1:]
	mut.Unlock()
	return cust
}

func printStatus() {
	fmt.Println("****************************************")
	fmt.Println("* Booked Tables: ", bookedTables)
	fmt.Println("* Free Tables: ", availableTables)
	fmt.Println("* Customers in waiting: ", customerQueue)
	fmt.Println("****************************************")
}
