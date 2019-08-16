package main

import "fmt"

func main() {
	words, err := fetchWords()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", words)
}
