package finder

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

type PrimeFinder struct {
	Ranges []PrimeFinderRange
}

type PrimeFinderRange struct {
	Start int
	End   int
}

func (f PrimeFinder) Execute() []int {
	var wg sync.WaitGroup

	mutex := sync.RWMutex{}
	primesResult := make([]int, 0)
	completedRanges := 0

	for _, finderRange := range f.Ranges {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var primes = finderRange.getPrimes()

			mutex.Lock()
			primesResult = append(primesResult, primes...)
			completedRanges++
			mutex.Unlock()
		}()
	}

	go func() {
		ticker := time.NewTicker(time.Second)

		for range ticker.C {
			mutex.RLock()
			fmt.Printf("Go ranges completion count: %d/%d (%.2f%%)\n", completedRanges, len(f.Ranges), float64(completedRanges)/float64(len(f.Ranges))*100)

			if completedRanges == len(f.Ranges) {
				fmt.Println("All ranges completed")
				ticker.Stop()
			}
			mutex.RUnlock()
		}
	}()

	wg.Wait()
	sort.Ints(primesResult)
	return primesResult
}

func (f *PrimeFinder) SetRange(start int, end int, numberPerBatch int) {
	primeFinderRanges := make([]PrimeFinderRange, 0)

	for i := start; i <= end; i += numberPerBatch {
		endNumber := i + numberPerBatch - 1
		if endNumber > end {
			endNumber = end
		}
		primeFinderRanges = append(primeFinderRanges, PrimeFinderRange{i, endNumber})
	}

	f.Ranges = primeFinderRanges
}

func (r PrimeFinderRange) getPrimes() []int {
	var primeNumbers []int
	for i := r.Start; i <= r.End; i++ {
		if isPrime(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
