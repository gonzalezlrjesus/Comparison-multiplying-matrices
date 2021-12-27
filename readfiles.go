package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadFiles read matrices data from a file
func ReadFiles(nameFile string, sizeMatrix int) ([][]int, [][]int, [][]int) {
	matrixA := make([][]int, sizeMatrix)
	for key := range matrixA {
		matrixA[key] = make([]int, sizeMatrix)
	}

	matrixB := make([][]int, sizeMatrix)
	for key := range matrixB {
		matrixB[key] = make([]int, sizeMatrix)
	}

	matrixSolution := make([][]int, sizeMatrix)
	for key := range matrixSolution {
		matrixSolution[key] = make([]int, sizeMatrix)
	}

	file, err := os.Open(nameFile)
	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrixReading := ""
	contRow := 0

	const MatrixA, MatrixB, Solution = "Matrix-A", "Matrix-B", "Solution"

	for scanner.Scan() {
		if scanner.Text() == MatrixA {
			matrixReading = MatrixA

			continue
		}

		if scanner.Text() == MatrixB {
			matrixReading = MatrixB
			contRow = 0

			continue
		}

		if scanner.Text() == Solution {
			matrixReading = Solution
			contRow = 0

			continue
		}

		switch matrixReading {
		case MatrixA:
			for ind, value := range strings.Fields(scanner.Text()) {
				parsedValue, errParseInt := strconv.Atoi(value)
				if errParseInt != nil {
					log.Panic(err)
				}

				matrixA[contRow][ind] = parsedValue
			}

			contRow++
		case MatrixB:
			for ind, value := range strings.Fields(scanner.Text()) {
				parsedValue, errParseInt := strconv.Atoi(value)
				if errParseInt != nil {
					log.Panic(err)
				}

				matrixB[contRow][ind] = parsedValue
			}

			contRow++
		case Solution:
			for ind, value := range strings.Fields(scanner.Text()) {
				parsedValue, errParseInt := strconv.Atoi(value)
				if errParseInt != nil {
					log.Panic(err)
				}

				matrixSolution[contRow][ind] = parsedValue
			}

			contRow++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return matrixA, matrixB, matrixSolution
}
