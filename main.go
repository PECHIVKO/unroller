package main

import "fmt"

// var input [][]int = [][]int{
// 	{1, 2, 3},
// 	{4, 5, 6},
// 	{7, 8, 9},
// 	{10, 11, 12},
// }

func unroll(rolled [][]int, unrolled []int) ([][]int, []int) {
	if len(rolled) > 0 {
		unrolled = append(unrolled, rolled[0]...)
		rolled = rolled[1:]
		rolled = rotate(rolled)
		rolled, unrolled = unroll(rolled, unrolled)
	}
	return rolled, unrolled
}

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

func makeSlice(m, n int) [][]int {
	slice := make([][]int, m)
	for i := range slice {
		slice[i] = make([]int, n)
	}
	return slice
}

func getInput() [][]int {
	var m, n int
	fmt.Print("Input dimmentions of [m][n]slise:\nm = ")
	fmt.Scan(&m)
	fmt.Printf("n = ")
	fmt.Scan(&n)
	if m != n {
		fmt.Println("Is not [n][n]slice but still can be unrolled")
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
	return newInput
}

func main() {
	var input [][]int
	var output []int
	input = getInput()
	_, output = unroll(input, output)
	fmt.Println("\nUnrollred input:\n", output)
}
