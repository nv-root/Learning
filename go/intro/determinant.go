package main

import "fmt"

func main() {

	var matrix [3][3]int

	for i := range 3 {
		for j := range 3 {
			fmt.Printf("Enter element in %d row, %d col: ", i+1, j+1)
			fmt.Scan(&matrix[i][j])
		}
	}

	det := matrix[0][0]*(matrix[1][1]*matrix[2][2]-matrix[1][2]*matrix[2][1]) - matrix[0][1]*(matrix[1][0]*matrix[2][2]-matrix[1][2]*matrix[2][0]) + matrix[0][2]*(matrix[1][0]*matrix[2][1]-matrix[1][1]*matrix[2][0])

	fmt.Println("Determinant:", det)

}

// 00 01 02
// 10 11 12
// 20 21 22
