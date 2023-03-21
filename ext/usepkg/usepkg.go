package main

import (
	"ext/usepkg/custompkg" // 1 모듈 내 패키지
	"fmt"

	"github.com/guptarohit/asciigraph" // 3 외부 저장소 패키지
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()
	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

// This is custom Package
// This is Github expkg Sample
//  10.00 ┤        ╭╮
//   9.00 ┤   ╭╮   ││
//   8.00 ┤   ││ ╭╮││
//   7.00 ┤   │╰╮││││╭╮
//   6.00 ┤  ╭╯ │││││││ ╭
//   5.00 ┤ ╭╯  ╰╯╰╯│││╭╯
//   4.00 ┤╭╯       ││││
//   3.00 ┼╯        ││││
//   2.00 ┤         ╰╯╰╯
