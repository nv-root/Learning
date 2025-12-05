package main

import "fmt"
import "errors"

type customError struct {
	code    int
	message string
	err     error
}

// customError type implements error interface
func (e *customError) Error() string {
	return fmt.Sprintf("ERROR %d : %s :: %v\n", e.code, e.message, e.err)
}

func doSomething() error {
	// return &customError{
	// 	code:    500,
	// 	message: "Something went wrong!",
	// }

	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			err:     err,
		}
	}

	return nil
}

func doSomethingElse() error {
	return errors.New("Internal error")
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Println(err)
		return
	}
}
