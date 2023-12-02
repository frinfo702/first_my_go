package main

import(
	
	"fmt"
)
func main(){
	l := []int{100, 300, 23, 11, 2, 4, 6, 4}
	var min int = l[0]
	for i := 0; i < 8; i++{
		if min >= l[i]{
			min = l[i]
		}
	}

	fmt.Println(min)

	m := map[string]int{
		"apple": 200, 
		"banana": 300,
		"grapes":150,
		"orange": 80,
		"papaya": 500,
		"kiwi": 90,
	}

	sum := 0
	for _, price := range m {
		sum += price
	}
	fmt.Println(sum)
}


