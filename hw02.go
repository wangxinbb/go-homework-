package main

import (
	"fmt"
)

//*生成
func xingxing() {
	//五角星的上面一个角
	for i1 := 1; i1 < 6; i1++ {
		for j1 := 1; j1 < 19-i1; j1++ {
			fmt.Print(" ")
		}
		for k1 := 1; k1 <= 2*i1-1; k1++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	//五角星的中间两个角
	for i2 := 1; i2 < 5; i2++ {
		for j2 := 1; j2 < 3*i2-3; j2++ {
			fmt.Print(" ")
		}
		for k2 := 1; k2 <= 42-6*i2; k2++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	//中间与下部相接的部分
	for i3 := 1; i3 < 3; i3++ {
		for j3 := 1; j3 < 12-i3; j3++ {
			fmt.Print(" ")
		}
		for k3 := 1; k3 <= 12+2*i3; k3++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	//五角星的下面两个角
	for i4 := 1; i4 < 5; i4++ {
		for j4 := 1; j4 < 10-i4; j4++ {
			fmt.Print(" ")
		}
		for k4 := 1; k4 <= 10-2*i4; k4++ {
			fmt.Print("*")
		}
		for m4 := 1; m4 < 6*i4-3; m4++ {
			fmt.Print(" ")
		}
		for n4 := 1; n4 <= 10-2*i4; n4++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

/*
//graph生成
func graphxing(){

}
*/
//作业2 用 Go 程序画一颗五角星，使用graph库
func main() {
	xingxing()
	//graphxing()
}
