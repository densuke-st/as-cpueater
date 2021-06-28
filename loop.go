package main

import (
	"fmt"
	"os"
)

func main() {
	file := "stop"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	fmt.Println("check file: " + file)
	min := -2147483648
	max := 2147483647
	count := 0
	for {
		for i := min; i < max; i++ {
			// do nothing
		}
		fmt.Println(count)
		count = count + 1
		_, err := os.Stat(file)
		if err == nil {
			break
		}
	}
	fmt.Println("OK")
}
