package main

import "fmt"

func isDesirabelNumber(number int) bool {
	return number%3 == 0 || number%5 == 0
}

func main() {

	for iterator := 1; iterator < 51; iterator++ {
		if isDesirabelNumber(iterator) {
			fmt.Println(iterator)
		}
	}

}
