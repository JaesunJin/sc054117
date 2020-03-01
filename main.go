package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLenght, uppercase := lenAndUpper("jason")
	fmt.Println(totalLenght, uppercase)
	repeatMe("jason", "jin")
}
