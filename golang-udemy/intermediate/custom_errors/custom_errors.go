package main

import (
	"errors"
	"fmt"
)

func main() {

	// if err := doSomething(); err!= nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Operation completed successfully")

	if err1 := doSomething1(); err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("completed successfully")
}

type customError struct {
	code int
	message string
}

func (e *customError) Error() string{
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func doSomething() error {
	return &customError{
		code: 500,
		message: "Something went wrong!",
	}
}


type customError1 struct{
	code int
	message string
	er error
}

func (e *customError1) Error() string {
	return fmt.Sprintf("Error %d: %s\nDetails: %v\n", e.code, e.message,e.er)
}

func doSomethingElse1() error{
	return errors.New("internal error")
}

func doSomething1() error {
	err := doSomethingElse1()
	if err != nil {
		return &customError1{
			code: 500,
			message: "Something went wrong",
			er : err,
		}
	}
	return nil
}