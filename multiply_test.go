package main

import (
	"fmt"
	"reflect"
	"testing"
)

// ---- Benchmark to size Matrix: 2 ----

func BenchmarkMultiplyWithoutConcurrency_Size2(b *testing.B) {
	const sizematrix int = 2
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyWithoutConcurrency(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerRow_Size2(b *testing.B) {
	const sizematrix int = 2
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerRow(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerElement_Size2(b *testing.B) {
	const sizematrix int = 2
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerElement(sizematrix, matrixA, matrixB, &matrixC)
	}
}

// ---- Benchmark to size Matrix: 10 ----

func BenchmarkMultiplyWithoutConcurrency_Size10(b *testing.B) {
	const sizematrix int = 10
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyWithoutConcurrency(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerRow_Size10(b *testing.B) {
	const sizematrix int = 10
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerRow(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerElement_Size10(b *testing.B) {
	const sizematrix int = 10
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerElement(sizematrix, matrixA, matrixB, &matrixC)
	}
}

// ---- Benchmark to size Matrix: 50 ----

func BenchmarkMultiplyWithoutConcurrency_Size50(b *testing.B) {
	const sizematrix int = 50
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyWithoutConcurrency(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerRow_Size50(b *testing.B) {
	const sizematrix int = 50
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerRow(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerElement_Size50(b *testing.B) {
	const sizematrix int = 50
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerElement(sizematrix, matrixA, matrixB, &matrixC)
	}
}

// ---- Benchmark to size Matrix: 100 ----

func BenchmarkMultiplyWithoutConcurrency_Size100(b *testing.B) {
	const sizematrix int = 100
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyWithoutConcurrency(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerRow_Size100(b *testing.B) {
	const sizematrix int = 100
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerRow(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerElement_Size100(b *testing.B) {
	const sizematrix int = 100
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerElement(sizematrix, matrixA, matrixB, &matrixC)
	}
}

// ---- Benchmark to size Matrix: 2000 ----
func BenchmarkMultiplyWithoutConcurrency_Size2000(b *testing.B) {
	const sizematrix int = 2000
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyWithoutConcurrency(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerRow_Size2000(b *testing.B) {
	const sizematrix int = 2000
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerRow(sizematrix, matrixA, matrixB, &matrixC)
	}
}

func BenchmarkMultiplyConcurrencyPerElement_Size2000(b *testing.B) {
	const sizematrix int = 2000
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)
	matrixA, matrixB, _ := ReadFiles(fileName, sizematrix)
	matrixC := make([][]int, sizematrix)

	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	for i := 0; i < b.N; i++ {
		MultiplyConcurrencyPerElement(sizematrix, matrixA, matrixB, &matrixC)
	}
}

// ---- Tests only for size Matrix: 2000 ----
func TestMultiplyWithoutConcurrency(t *testing.T) {
	const sizematrix int = 2000
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)

	matrixA, matrixB, matrixSolution := ReadFiles(fileName, sizematrix)

	matrixC := make([][]int, sizematrix)
	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	MultiplyWithoutConcurrency(sizematrix, matrixA, matrixB, &matrixC)

	for indRow := 0; indRow < sizematrix; indRow++ {
		if !reflect.DeepEqual(matrixSolution[indRow], matrixC[indRow]) {
			t.Errorf("expected matrixSolution not equal to got MatrixC")
		}
	}
}

func TestMultiplyConcurrencyPerRow(t *testing.T) {
	const sizematrix int = 2000
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)

	matrixA, matrixB, matrixSolution := ReadFiles(fileName, sizematrix)

	matrixC := make([][]int, sizematrix)
	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	MultiplyConcurrencyPerRow(sizematrix, matrixA, matrixB, &matrixC)

	for indRow := 0; indRow < sizematrix; indRow++ {
		if !reflect.DeepEqual(matrixSolution[indRow], matrixC[indRow]) {
			t.Errorf("expected matrixSolution not equal to got MatrixC")
		}
	}
}

func TestMultiplyConcurrencyPerElement(t *testing.T) {
	const sizematrix int = 2000
	fileName := fmt.Sprintf("matrices%vx%v.txt", sizematrix, sizematrix)

	matrixA, matrixB, matrixSolution := ReadFiles(fileName, sizematrix)

	matrixC := make([][]int, sizematrix)
	for key := range matrixC {
		matrixC[key] = make([]int, sizematrix)
	}

	MultiplyConcurrencyPerElement(sizematrix, matrixA, matrixB, &matrixC)

	for indRow := 0; indRow < sizematrix; indRow++ {
		if !reflect.DeepEqual(matrixSolution[indRow], matrixC[indRow]) {
			t.Errorf("expected matrixSolution not equal to got MatrixC")
		}
	}
}
