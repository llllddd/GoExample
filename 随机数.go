package main

import (
	"fmt"
	"math/rand"
	"time"
)

//使用时间作为种子

func main() {
	//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//	for i := 0; i < 100; i++ {
	//		fmt.Println(r.Intn(1000))
	//	}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10000; i++ {
		fmt.Println(rand.Intn(10000))
	}
}
