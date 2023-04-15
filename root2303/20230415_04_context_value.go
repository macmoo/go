package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	fmt.Println("---------------------------------")
	// ctx := context.WithValue(context.Background(), "number", 9)
	// OR
	ctx := context.Background()
	ctx = context.WithValue(ctx, "number", 9)
	ctx = context.WithValue(ctx, "Key", "Tulip")
	go square(ctx)
	wg.Wait()
}

func square(ctx context.Context) {
	// 컨텍스트에서 값을 읽기
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d\n", n*n)
	}
	if v := ctx.Value("Key"); v != nil {
		s := v.(string)
		fmt.Printf("Name:%s\n", s)
	}
	wg.Done()
}

// ---------------------------------
// Square:81
// Name:Tulip
// ---------------------------------
