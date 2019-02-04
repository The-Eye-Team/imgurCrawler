package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/labstack/gommon/color"
)

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]")
var tildPre = color.Yellow("[") + color.Green("~") + color.Yellow("]")
var crossPre = color.Yellow("[") + color.Red("✗") + color.Yellow("]")
var idRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateID(length int) string {
	id := make([]rune, length)
	for i := range id {
		id[i] = idRunes[rand.Intn(len(idRunes))]
	}
	return string(id)
}

func crawlID(index int, worker *sync.WaitGroup) {
	defer worker.Done()
	fmt.Print(index)
	fmt.Println(" - " + generateID(7))
}

func main() {
	var worker sync.WaitGroup
	var count, index int

	// Parse arguments
	parseArgs(os.Args)

	for {
		worker.Add(1)
		count++
		go crawlID(index, &worker)
		index++
		if count == arguments.Concurrency {
			worker.Wait()
			count = 0
		}
	}
}
