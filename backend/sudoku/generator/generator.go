package main

import (
	"C"
	"time"
	"math/rand"
)

var NUM_OF_SHUFFLES int = 100

func main() {}

//export generatePuzzleToPython
func generatePuzzleToPython(ins []int, sol []int) {
	rand.Seed(time.Now().UTC().UnixNano())
	instance := createSolvedSudoku()
	
	// write solution
	for i := 0; i < 9; i ++ {
		for j := 0; j < 9; j ++ {
			sol[i * 9 + j] = instance[i][j] 
		}
	}
	
	// write puzzle instance
	createdUnsolvedSudoku(instance)
	for i := 0; i < 9; i ++ {
		for j := 0; j < 9; j ++ {
			ins[i * 9 + j] = instance[i][j] 
		}
	}
}

func createSolvedSudoku() [][]int {
	// Base instance
	instance := make([][]int, 9)
	instance[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	instance[1] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	instance[2] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	instance[3] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1}
	instance[4] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	instance[5] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	instance[6] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2}
	instance[7] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5}
	instance[8] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8}

	// Shuffle the generated Sudoku instance with different methods
	for i := 0; i < NUM_OF_SHUFFLES; i ++ {
		var random int = rand.Intn(3)
		if random == 0 {
			swapRow(instance)
		} else if random == 1 {
			swapCol(instance)
		} else {
			swapNumber(instance)
		}
	}
	return instance
}

// swap 2 random rows in a 9x3 block within the Sudoku instance
func swapRow(grid [][]int) {
	var row int = (rand.Intn(3) / 3) * 3
	var l1, l2 int = row + rand.Intn(3), row + rand.Intn(3)
	for l1 == l2{
		l2 = row + rand.Intn(3)
	}
	grid[l1], grid[l2] = grid[l2], grid[l1]
}

// swap 2 random columns in a 3x9 block within the Sudoku instance
func swapCol(grid [][]int) {
	var col int = (rand.Intn(3) / 3) * 3
	var c1, c2 int = col + rand.Intn(3), col + rand.Intn(3)
	for c1 == c2{
		c2 = col + rand.Intn(3)
	}

	for line := 0; line < 9; line++ {
		grid[line][c1], grid[line][c2] = grid[line][c2], grid[line][c1]
	}
}

// Swap positions of 2 random values in the Sudoku instance
func swapNumber(grid [][]int) {
	var n1, n2 int = 1 + rand.Intn(8), 1 + rand.Intn(8)
	for n1 == n2 {
		n2 = 1 + rand.Intn(8)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == n1 {
				grid[i][j] = n2
			} else if grid[i][j] == n2 {
				grid[i][j] = n1
			}
		}
	}
}

// Removing numbers from the Sudoku instance, while making sure
// that the puzzle only has one unique solution
func createdUnsolvedSudoku(grid [][]int) {
	digits := make([]int, 81)
	for i := range digits {
		digits[i] = i
	}
	rand.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i] 
	})

	// Randomly removes a number, and check if there still is an 
	// unique solution. If not, undo the move. Do this for every 
	// number in the puzzle 
	for i := range digits {
		var row int = digits[i] / 9
		var col int = digits[i] % 9
		temp := grid[row][col]
		grid[row][col] = 0

		if !hasUniqueSolution(grid) {
			grid[row][col] = temp
		}
	}
}

// Check if there is an unique solution by running the solver 
// twice and return false if a new solution is found. 
func hasUniqueSolution(grid [][]int) bool{
	var arr [9][9]int
	var arr2 [9][9]int
	for i := range arr {
		for j := range arr[i] { 
			arr[i][j] = grid[i][j]
			arr2[i][j] = grid[i][j]
		}
	}
	
	if backtrack(&arr, &arr2) {
		if backtrack(&arr2, &arr) {
			return false
		}
		return true
	} else {
		return false
	}
}

// solver which uses simple backtracking algorithm
func backtrack(board *[9][9]int, constraint *[9][9]int) bool {
	if !hasEmptyCell(board) {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != constraint[i][j] {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board, constraint) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {
	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}