package main

import (
	"fmt"
	"os"
	"prime/constants"
	"prime/finder"
	"strconv"
	"time"
)

func main() {
	var min int = 1
	var max int = constants.Giga

	args := os.Args
	if len(args) == 3 {
		minArg := args[1]
		maxArg := args[2]

		if x, err := strconv.Atoi(minArg); err == nil {
			min = x
		} else {
			fmt.Println("Invalid min argument")
			return
		}

		if y, err := strconv.Atoi(maxArg); err == nil {
			max = y
		} else {
			fmt.Println("Invalid max argument")
			return
		}

		isValidMinMax := validateMinMax(min, max)
		if !isValidMinMax {
			return
		}
	}

	fmt.Printf("Finding prime numbers from %d to %d\n", min, max)

	startTime := time.Now()

	primeFinder := finder.PrimeFinder{}
	primeFinder.SetRange(min, max, 1000)
	finder.WritePrimeFinderRangesToFile(primeFinder.Ranges)

	primeNumbers := primeFinder.Execute()
	finder.WritePrimesToFile(primeNumbers)

	endTime := time.Now()

	numberPerSecond := float64(len(primeNumbers)) / endTime.Sub(startTime).Seconds()
	fmt.Printf("Found prime numbers from %d to %d\n", min, max)
	fmt.Printf("Execution time: %v\n", endTime.Sub(startTime))
	fmt.Printf("Number of prime numbers: %d\n", len(primeNumbers))
	fmt.Printf("Number found per second: %.2f\n", numberPerSecond)
}

func validateMinMax(min int, max int) bool {
	if min < 1 {
		fmt.Println("Min cannot be less than 1")
		return false
	}

	if max < 1 {
		fmt.Println("Max cannot be less than 1")
		return false
	}

	if min == max {
		fmt.Println("Min and max cannot be the same")
		return false
	}

	if min > max {
		fmt.Println("Min cannot be greater than max")
		return false
	}

	return true
}
