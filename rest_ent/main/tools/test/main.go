package main

import (
	"fmt"
	"time"
)

func test_panic() error{
	panic("abcde")
	fmt.Println("zzzz")
	return nil
}

type Test struct {
	createdAt time.Time
}

func main() {
	fmt.Println(Test{}.createdAt);


}