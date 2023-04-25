package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var availableTables = []int{}
var bookedTables = []int{}
var cQueue []int
var queueLen = 0
var custCount = 0
var maxOrders = 15
var minSec = 1
var maxSec = 5
var minDiningPeriod = 10
var maxDiningPeriod = 15

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := getInput("Please enter no of tables: ", reader)
	q, _ := getInput("Please enter customer queue length: ", reader)
	queueLen = parseInt(q)
	// fmt.Print("hello: ")
	// fmt.Scanf("%d", &queueLen)

	// fmt.Printf("Queue: %v \n", queueLen)
	rand.Seed(time.Now().UnixNano())
	createTables(parseInt(t))
	startCustomerGeneration()
}
