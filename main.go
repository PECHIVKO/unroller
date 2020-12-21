package main

import "fmt"

func unroll(rolled [][]int, unrolled []int) ([][]int, []int) {
	if len(rolled) > 0 {
		// writing rolled[0] to output slice
		unrolled = append(unrolled, rolled[0]...)
		// deleting rolled[0] from rolled
		rolled = rolled[1:]
		// rotating rolled counterclockwise
		rolled = rotate(rolled)
		// recursion
		rolled, unrolled = unroll(rolled, unrolled)
	}
	return rolled, unrolled
}

//rotate rotates input [][]int slice counterclockwise
func rotate(unrotated [][]int) [][]int {
	if len(unrotated) > 0 {
		rotated := makeSlice(len(unrotated[0]), len(unrotated))

		for i := range unrotated {
			for j := range unrotated[i] {
				rotated[len(rotated)-1-j][i] = unrotated[i][j]
			}
		}
		return rotated
	}
	return unrotated
}

// makeSlice creates slice with m*n dimmentions
func makeSlice(m, n int) [][]int {
	slice := make([][]int, m)
	for i := range slice {
		slice[i] = make([]int, n)
	}
	return slice
}

// createMatrix gets user input, creates Matrix with user values
func createMatrix() ([][]int, error) {
	var m, n int
	var err error
	fmt.Print("Input dimmentions of [m(rows)][n(columns)]slise:\nm = ")
	fmt.Scan(&m)
	fmt.Printf("n = ")
	fmt.Scan(&n)
	if m != n {
		fmt.Println("Is not [n][n]slice but still can be unrolled")
		// OR
		// err = fmt.Errorf("Not square [m][n]slice IS NOT ALLOWED")
		// return nil, err
	}
	newInput := makeSlice(m, n)
	fmt.Println("Input every value in slice (You can use any single non-numeric character to separate values):")
	for i := range newInput {
		for j := range newInput[i] {
			fmt.Scan(&newInput[i][j])
			if j < len(newInput[i])-1 {
			}
		}
	}
	fmt.Println("\nYour input was:")
	for _, row := range newInput {
		fmt.Println(row)
	}
	return newInput, err
}

func main() {
	var input [][]int
	var output []int
	input, err := createMatrix()
	if err != nil {
		fmt.Println("err")
	} else {
		_, output = unroll(input, output)
		fmt.Println("\nUnrollred input:\n", output)
	}
}
