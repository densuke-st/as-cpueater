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
	for {
		_, err := os.Stat(file)
		if err == nil {
			break
		}
	}
	fmt.Println("OK")
}
