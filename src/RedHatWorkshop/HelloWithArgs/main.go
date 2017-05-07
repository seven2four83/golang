package main

import "fmt"
import "os"
import "strconv"


func isOdd(value int) bool{
	return value % 2 != 0
}

func main() {
	// fmt.Println("Hello, 7to4 is a noob Gopher!")
	argument, error2 :=  strconv.Atoi(os.Args[1])
	if error2 == nil{
		fmt.Println (isOdd(argument))
	}else{
		fmt.Println ("Error converting")
	}
}
