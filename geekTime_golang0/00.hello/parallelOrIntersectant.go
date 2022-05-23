package main

import (
	"fmt"
	//	_ "github.com/armstrongli/learn.go"
)

func main() {
	var a [2]float64
	fmt.Println("输入第一坐标点: ")
	fmt.Scanln(&a[0], &a[1])
	var b [2]float64
	fmt.Println("输入第二坐标点: ")
	fmt.Scanln(&b[0], &b[1])
	var c = a[0] - b[0]
	var d = a[1] - b[1]
	var e = c / d
	var f [2]float64
	fmt.Println("输入第三坐标点: ")
	fmt.Scanln(&f[0], &f[1])
	var g [2]float64
	fmt.Println("输入第四坐标点: ")
	fmt.Scanln(&g[0], &g[1])
	var h = f[0] - g[0]
	var i = f[1] - g[1]
	var j = h / i
	if e/j == 1 {
		fmt.Println("两条直线平行")
	} else {
		fmt.Println("两条直线不平行")
	}

}
