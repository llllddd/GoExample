package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe   bool       = false
	Maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, Tobe, Tobe)
	fmt.Printf(f, Maxint, Maxint)
	fmt.Printf(f, z, z)
	//%T显示变量类型,%v显示变量的值.
}
