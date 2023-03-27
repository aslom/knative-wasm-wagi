package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Content-Type: text/plain")
	fmt.Println("")
	fmt.Println("Hello After Pi Day in Wasm World!")

	printenv()
}

func printenv() {
	//fmt.Printf("HOME=%s\n", os.Getenv("HOME"))
	//fmt.Println(os.Environ())
	for _, e := range os.Environ() {
		fmt.Printf("%s\n", e)
	}
}
