package main

import "fmt"

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}


func findWinner(n, m int) int {
	var inside = n
	var curPos int
	onesSlice := make([]int, n)
	for i:= 0; i < len(onesSlice); i++{
		onesSlice[i] = i+1
	}
	for n !=0 {
		curPos = curPos+m -1
		for curPos >=inside {
			curPos = curPos-inside
		}
		onesSlice = remove(onesSlice,curPos)
		inside = inside - 1
		if len(onesSlice)==1 {
			return onesSlice[0]
		}

	}
	return 0
}

func main() {
	var n,m  int
	fmt.Println("Enter n:")
	fmt.Scanln(&n)
	fmt.Println("Enter m:")
	fmt.Scanln(&m)

	fmt.Printf("Vuvedohte za n -> %d , m ->  %d ",
		n, m)

	fmt.Println("winner is : " , findWinner(n,m))
}