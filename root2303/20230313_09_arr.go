package main

import "fmt"

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	for i, arr := range a {
		for j, v := range arr {
			fmt.Printf("a[%d][%d]=%d\n", i, j, v)
		}
	}
}

// a[0][0]=1
// a[0][1]=2
// a[0][2]=3
// a[0][3]=4
// a[0][4]=5
// a[1][0]=6
// a[1][1]=7
// a[1][2]=8
// a[1][3]=9
// a[1][4]=10
