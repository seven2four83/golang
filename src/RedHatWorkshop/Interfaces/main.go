package main

import "fmt"
import "os"
import "strconv"

func isDesirabelNumber(number int) bool {
	return number%3 == 0 || number%5 == 0
}


type Shape interface{
	area
}

type Rectangle struct{
	length float64
	width float64
} 

func (rect Rectangle) area() float64{
	return rect.length * rect.width
}

func getFloat64(inputString string) float64{
	floatValue, _ := strconv.ParseFloat(inputString,2)
	return floatValue
}

func main() {
	if os.Args[1] == "c"{

	}else if os.Args[1] == "r"{
		var rect Rectangle 
		rect.length = getFloat64(os.Args[2])
		rect.width = getFloat64(os.Args[3])
		fmt.Println(rect.length,rect.width)
		fmt.Println(rect.area())
	}
}
