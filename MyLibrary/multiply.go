package calculate

import "fmt"

var Percent int = 10;

type deck []string

func LoopArrayByType(d deck){
	for i, deckItem := range d {
		fmt.Println(i,deckItem)
	}
}

func MyMultiply(x int, y int) int {
	total := x * y
	return total
}

func LoopByValue(count int) {

	for i := 0; i < count; i++ {
		fmt.Println(i)
	}	
}

func LoopArray(arr []string){
	for i, arrItem := range arr{
		fmt.Println(i,arrItem)
	}
}

func FuncFloat() float64{
	return 4.44
}
