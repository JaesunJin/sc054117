package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("i'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func arr() {
	names := []string{"j", "a"}
	names = append(names, "s")
	fmt.Println(names)
}

func learnMap() {
	jason := map[string]string{"name": "jason", "age": "30"}
	for key, value := range jason {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func learnStructure() {
	favFood := []string{"gyoza", "potato"}
	jason := person{name: "jason", age: 32, favFood: favFood}
	fmt.Println(jason)
}

func main() {
	// learnMap()
	learnStructure()
}
