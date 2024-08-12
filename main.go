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
	var minNum = 1
	var maxNum = 100 * constants.Mil

	args := os.Args
	if len(args) == 3 {
		minArg := args[1]
		maxArg := args[2]

		if x, err := strconv.Atoi(minArg); err == nil {
			minNum = x
		} else {
			fmt.Println("Invalid minNum argument")
			return
		}

		if y, err := strconv.Atoi(maxArg); err == nil {
			maxNum = y
		} else {
			fmt.Println("Invalid maxNum argument")
			return
		}

		isValidMinMax := validateMinMax(minNum, maxNum)
		if !isValidMinMax {
			return
		}
	}

	fmt.Printf("Finding prime numbers from %d to %d\n", minNum, maxNum)

	startTime := time.Now()

	primeFinder := finder.PrimeFinder{}
	primeFinder.SetRange(minNum, maxNum, 1000)
	finder.WritePrimeFinderRangesToFile(constants.PrimeRangesFilePath, primeFinder.Ranges)

	primeNumbers := primeFinder.Execute()
	finder.WritePrimesToFile(constants.PrimeFilePath, primeNumbers)

	endTime := time.Now()

	numberPerSecond := float64(len(primeNumbers)) / endTime.Sub(startTime).Seconds()
	fmt.Printf("Found %d prime numbers in the range %d to %d\n", len(primeNumbers), minNum, maxNum)
	fmt.Printf("Execution time: %v\n", endTime.Sub(startTime))
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
