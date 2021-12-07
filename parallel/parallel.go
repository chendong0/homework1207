package main

import "fmt"

func main() {
	/*
		平面内两条线平行条件:
			作业：判断两条线是否平行
		提示1.两点决定一条直线
			2.两条线是否平行取决于两条线的斜率是否一样
	*/
	//变量标准声明
	var line1 float64
	//变量批量声明
	var (
		//line1 float64
		y1 float64
		k  float64
		x  float64
		b  float64
	)
	//变量批量声明
	var (
		line2 float64
		y2    float64
		c     float64
	)
	//给y1 y2赋值
	y1 = ((k * x) + b)
	y2 = ((k * x) + c)

	//给line1 line2初始化多个变量
	line1, line2 = y1, y2

	if line1 == line2 {
		fmt.Println("两条直线斜率相等，结论为两条直线平行")
	} else if line1 != line2 {
		fmt.Println("两条直线斜率不相等，结论两条直线不平行")
	}

}
