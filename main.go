package main

import (
	"fmt"
	"sync"
)

func main() {
	testSlice := []int{24, 10, 4, 16, 20, 9, 26, 18, 12, 27, 19, 22, 14, 11, 23, 1, 2, 30, 5, 8, 28, 31, 29, 6, 3, 17, 13, 7, 15, 25, 32, 21}
	bitonicSort(testSlice, 1)
	fmt.Println(testSlice)
}

func bitonicSort(elements []int, dir int) {
	l := len(elements)

	if !isPowerOfTwo(l) {
		fmt.Println("input is not power of two")
		return
	}

	if l > 1 {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			bitonicSort((elements)[:l/2], 1)
		}()

		go func() {
			defer wg.Done()
			bitonicSort((elements)[l/2:], 0)
		}()

		wg.Wait()
		bitonicMerge(elements, dir)
	}
}

func bitonicMerge(elements []int, dir int) {
	l := len(elements)
	if l > 1 {
		mid := l / 2
		for i := 0; i < mid; i++ {
			compareAndSwap(elements, i, i+mid, dir)
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			bitonicMerge(elements[:mid], dir)
		}()

		go func() {
			defer wg.Done()
			bitonicMerge(elements[mid:], dir)
		}()

		wg.Wait()
	}
}

func compareAndSwap(s []int, i, j, dir int) {
	if s[i] > s[j] && dir == 1 || s[i] < s[j] && dir == 0 {
		s[i], s[j] = s[j], s[i]
	}
}

func isPowerOfTwo(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}
