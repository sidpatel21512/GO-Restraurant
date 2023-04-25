package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	wg := &sync.WaitGroup{}
	customerC := make(chan int)
	go startRestaurant(wg)
	go generateCustomers(wg, customerC)
	wg.Add(1)
	go startTableAssignment(wg, customerC)

	wg.Wait()
	fmt.Println("We are closed!!!")
}

func startRestaurant(wg *sync.WaitGroup) {
	fmt.Println("We are now open...")
	for range time.Tick(time.Second * 1) {
		if restaurantTiming == 0 {
			mut.Lock()
			isRestaurantTimeOver = true
			mut.Unlock()
			break
		}
		restaurantTiming--
	}
	wg.Add(1)
	wg.Done()
}

func startTableAssignment(wg *sync.WaitGroup, c chan int) {
	for {
		_, isChannelOpened := <-c

		mut.Lock()
		if isRestaurantTimeOver && isRestaurantClosed {
			break
		}
		mut.Unlock()

		if isChannelOpened {
			cust := removeCustomer()

			mut.Lock()
			tableCount := len(availableTables)
			mut.Unlock()
			if tableCount > 1 {
				go assignTable(wg, c, tableCount, &cust)
			} else if tableCount == 1 {
				go assignTable(wg, c, 2, &cust)
			} else {
				fmt.Printf("Customer %v has left the restaurant, due to waiting. \n", cust.id)
			}
		}
	}

	defer wg.Done()
}

func assignTable(wg *sync.WaitGroup, ch chan int, count int, c *Customer) {
	no := getRandomNumber(0, count-1)
	c.tableNo = availableTables[no]
	bookTable(c.tableNo)
	mut.Lock()
	availableTables = append(availableTables[:no], availableTables[no+1:]...)
	mut.Unlock()
	fmt.Printf("Customer %v has been assigned a table %v. \n", c.id, c.tableNo)

	time.AfterFunc(time.Duration(c.dinnerTime)*time.Second, func() {
		freeTheTable(c.tableNo)
		checkRestaurantClosed(ch)
		defer wg.Done()
	})
	wg.Add(1)
}

func freeTheTable(t int) {
	mut.Lock()
	i := indexOf(bookedTables, t)
	bookedTables = append(bookedTables[:i], bookedTables[i+1:]...)
	availableTables = append(availableTables, t)
	mut.Unlock()
	fmt.Printf("Table %v is available. \n", t)
}

func bookTable(n int) {
	mut.Lock()
	bookedTables = append(bookedTables, n)
	mut.Unlock()
}

func checkRestaurantClosed(c chan int) {
	mut.Lock()
	if isRestaurantTimeOver && len(bookedTables) == 0 {
		isRestaurantClosed = true
		c <- 0
		close(c)
	}
	mut.Unlock()
}
