package main

import "fmt"

func main() {
	for {
		go fmt.Println(0)
		fmt.Print(1)
	}

}
