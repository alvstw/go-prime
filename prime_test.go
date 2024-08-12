package main

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkPrimeFinder(b *testing.B) {
	startTime := time.Now()

	finder := PrimeFinder{}

	finder.setRange(1, 1000000, 100)
	writePrimeFinderRangesToFile(finder.Ranges)

	primeNumbers := finder.execute()
	writePrimesToFile(primeNumbers)

	endTime := time.Now()

	numberPerSecond := float64(len(primeNumbers)) / endTime.Sub(startTime).Seconds()
	fmt.Printf("Execution time: %v\n", endTime.Sub(startTime))
	fmt.Printf("Number of prime numbers: %d\n", len(primeNumbers))
	fmt.Printf("Number found per second: %.2f\n", numberPerSecond)
}
