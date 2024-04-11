package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	numbers := generateConsecutiveNumbers(1000)
	//no goroutin
	startTimeSequential := time.Now()
	processNumbersSequentially(numbers)
	endTimeSequential := time.Now()
	fmt.Printf("Execution time without goroutines: %v\n", endTimeSequential.Sub(startTimeSequential))

	//goroutin
	numGoroutines := []int{1, 10, 100, 1000}
	for _, num := range numGoroutines {
		fmt.Printf("Processing with %d goroutines:\n", num)
		processNumbers(numbers, num)
	}
}

func generateConsecutiveNumbers(n int) []float64 {
	numbers := make([]float64, n)
	for i := 0; i < n; i++ {
		numbers[i] = float64(i)
	}
	return numbers
}

func processNumbersSequentially(numbers []float64) { //process for no goroutine
	for _, num := range numbers {
		_ = math.Sqrt(num)
	}
}

func processNumbers(numbers []float64, numGoroutines int) { //process for goroutine
	startTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	chunkSize := len(numbers) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				processNumber(numbers[j])
			}
		}(i*chunkSize, (i+1)*chunkSize)
	}

	wg.Wait()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time with %d goroutines: %v\n", numGoroutines, elapsedTime)
}

func processNumber(num float64) {
	_ = math.Sqrt(num)
}
