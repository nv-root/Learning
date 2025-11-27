package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println("Empty array:", a) // all 0s

	a[2] = 3
	fmt.Println("set:", a)
	fmt.Println("get:", a[2])

	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println("length of b:", len(b))
	fmt.Println("declared and initialized:", b)

	b = [...]int{11, 22, 33, 44, 55}
	fmt.Println(b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("zeroed:", b)

	var twoD [2][3]int
	for i:=range 2{
		for j:=range 3{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2D array:", twoD)

	// bubble sort
	var arr [5]int
	fmt.Println(arr)

	for i := range len(arr){
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr)

	for i := range len(arr)-1{
		for j := range len(arr)-1-i {
			if arr[j+1] < arr[j]{
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			} 
		}
	}

	fmt.Println("Sorted:", arr)

}
