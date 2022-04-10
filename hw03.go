package main

import (
	"flag"
	"fmt"
)

//作业3 计算圆面积并输出结果（要求有常量）
func main() {
	var R float64
	flag.Float64Var(&R, "r", 5, "输入半径 -r")
	flag.Parse()

	fmt.Println(sum(R))
}

func sum(r float64) float64 {
	var pi = 3.14
	sums := pi * r * r
	return sums
}
