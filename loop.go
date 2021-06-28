package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	file := "stop"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	fmt.Println("check file: " + file)
	min := -2147483648
	max := 2147483647

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
		s := <-sig
		fmt.Printf("Received signal: %s", s)
		os.Exit(0)
	}()

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
