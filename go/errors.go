package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Error: Can't compute square root for negative number")
	}

	return math.Pow(x, 0.5), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: empty")
	}
	return nil
}

// custom errors
type myError struct {
	message string
}

func (me *myError) Error() string {
	return fmt.Sprintf("Error: %s", me.message)
}

func eProcess() error {
	return &myError{"Custom error message"}
}

func readConfig() error {
	return errors.New("Config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}

	return nil
}

func main() {

	root, err := sqrt(-31)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(root)
	}

	data := []byte{}
	fmt.Println("Data:", data)

	if err := process(data); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data processed successfully")
	}

	if err_e := eProcess(); err_e != nil {
		fmt.Println(err_e)
	}

	if err_r := readData(); err_r != nil {
		fmt.Println(err_r)
	} else {
		fmt.Println("Data read successfully")
	}

}
