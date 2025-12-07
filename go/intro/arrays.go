package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println("Empty array:", a) // all 0s

	var names [3]string
	fmt.Println(names)
	names[2] = "three"
	names[1] = "two"
	fmt.Println(names)
	fmt.Println(len(names))

	originalArray := [2]int{1, 2}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", copiedArray)

	
	var arrPtr *[2]int
	arrPtr = &originalArray
	arrPtr[0] = 101

	fmt.Println("Original array:", originalArray)


	a[2] = 3
	fmt.Println("set:", a)
	fmt.Println("get:", a[2])

	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println("length of b:", len(b))
	fmt.Println("declared and initialized:", b)

	b = [...]int{11, 22, 33, 44, 55}
	fmt.Println(b)

	// looping over array
	for _, value := range b{
		fmt.Println(value)
	}

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
