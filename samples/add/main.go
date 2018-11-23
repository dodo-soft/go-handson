package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}
	s, err := add(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(s)
}

func add(s1 string, s2 string) (int, error) {
	num1, err := strconv.Atoi(s1)
	if err != nil {
		return -1, err
	}
	num2, err := strconv.Atoi(s2)
	if err != nil {
		return -1, err
	}
	return num1 + num2, nil
}
