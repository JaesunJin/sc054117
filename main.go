package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLenght, _ := lenAndUpper("jason")
	fmt.Println(totalLenght)
	repeatMe("jason", "jin")
}