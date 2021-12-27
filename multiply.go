package main

import (
	"sync"
)

// MultiplyWithoutConcurrency .
func MultiplyWithoutConcurrency(sizematrix int, matrixA [][]int, matrixB [][]int, matrixC *[][]int) {
	for i := 0; i < sizematrix; i++ {
		for j := 0; j < sizematrix; j++ {
			sum := 0
			for k := 0; k < sizematrix; k++ {
				sum += matrixA[i][k] * matrixB[k][j]
			}

			(*matrixC)[i][j] = sum
		}
	}
}

// MultiplyConcurrencyPerRow .
func MultiplyConcurrencyPerRow(sizematrix int, matrixA [][]int, matrixB [][]int, matrixC *[][]int) {
	var wg sync.WaitGroup

	for i := 0; i < sizematrix; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for j := 0; j < sizematrix; j++ {
				sum := 0
				for k := 0; k < sizematrix; k++ {
					sum += matrixA[i][k] * matrixB[k][j]
				}

				(*matrixC)[i][j] = sum
			}
		}(i)
	}
	wg.Wait()
}

// MultiplyConcurrencyPerElement .
func MultiplyConcurrencyPerElement(sizematrix int, matrixA [][]int, matrixB [][]int, matrixC *[][]int) {
	var wg sync.WaitGroup

	for i := 0; i < sizematrix; i++ {
		for j := 0; j < sizematrix; j++ {
			wg.Add(1)

			go func(i, j int) {
				defer wg.Done()

				sum := 0
				for k := 0; k < sizematrix; k++ {
					sum += matrixA[i][k] * matrixB[k][j]
				}

				(*matrixC)[i][j] = sum
			}(i, j)
		}
	}
	wg.Wait()
}
