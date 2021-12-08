package main

import (
	"fmt"
)

//全局变量声明
var totalFatRates float64

func main() {

	fmt.Println("hello1208") //启动信息

	//短变量声明
	Weights := [3]float64{} //体重(KG)

	//多变量声明

	Names, Talls, Ages, bmis := [3]string{}, [3]float64{}, [3]int{}, [3]float64{}
	//身高(m)
	//不能var 和 := 一起使用
	//name string := "chen"//将会报错,因为编译器会自动推断

	FatRates := [3]float64{}
	fmt.Print(FatRates) //测试FatRates

	for i := 0; i < 3; i++ {

		fmt.Print("\n姓名:")
		fmt.Scanln(&Names[i])

		fmt.Print("体重(KG):")
		fmt.Scanln(&Weights[i])

		fmt.Print("身高(米):")
		fmt.Scanln(&Talls[i])

		bmis[i] = (Weights[i] / (Talls[i] * Talls[i]))
		fmt.Printf("%s的bim指数为%.2f \n", Names, bmis)

		fmt.Print("年龄:")
		fmt.Scanln(&Ages[i]) //是读取到i上

		//单变量声明
		var SexWeight int
		var Sex string = "male"
		fmt.Println("性别(female/male):")
		fmt.Scanln(&Sex)

		if Sex == "male" { //==要区分理解=和==的区别和用法.
			SexWeight = 0
		} else {
			SexWeight = 1
		}

		//var (
		//	FatRates[i] = (1.2*bmis[i] + 0.23 * float64(Ages[i])-5.4 - 10.8 * float64(SexWeight)) / 100
		//)  //var( FatRates[i])会报错

		//var FatRates float64= ((1.2*bmis[i] + 0.23*float64(Ages[i]) - 5.4 - 10.8*float64(SexWeight)) / 100)
		var FatRate float64 = ((1.2*bmis[i] + 0.23*float64(Ages[i]) - 5.4 - 10.8*float64(SexWeight)) / 100)
		//var FatRates float64 应该改为var FatRates
		//FatRates[i] = FatRates

		FatRates[i] = FatRate //FatRates重名了,花费了半个小时.问同学才问出来.

		fmt.Printf("%s的体脂率为%.2f \n", Names, FatRate)

		fmt.Println("+++++---+++") //打印测试

		if Ages[i] < 18 {
			fmt.Println("不评估18岁以下,因为波动太大,效果没有意义")
		} //else if嵌套使用,熟料度,和练习相互嵌套
		//要把共有条件先放在前面
		//func calculator() {
		//	fmt.Println("+++++++++++-----------+++++++++++")
		//		}

		if Sex == "male" {
			//编写男性体脂率与体脂状态

			if Ages[i] >= 18 && Ages[i] <= 39 {
				if FatRate < 0.05 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:偏瘦,加强营养多锻炼\n", Names, bmis, FatRates)
					//%.2f 评估双引号里面的空格也会导致读取数据错误
				} else if FatRate >= 0.05 && FatRate < 0.1 { //5%报错,只能用0.15
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:您的状态偏瘦\n", Names, bmis, FatRates)
				} else if FatRate >= 0.1 && FatRate < 0.17 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:您目前标准,请保持\n", Names, bmis, FatRates)
				} else if FatRate >= 0.17 && FatRate < 0.22 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:您达到标准警戒线\n", Names, bmis, FatRates)

				} else if FatRate >= 0.22 && FatRate < 0.27 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:您轻度肥胖\n", Names, bmis, FatRates)
				} else if FatRate >= 0.27 && FatRate < 0.45 {
					fmt.Printf("bmis指数为:%.2f,体脂率为%.2f,评估建议:您重度肥胖\n", bmis, FatRates)
				} //对比作业,多加了左括号

			} else if Ages[i] > 39 && Ages[i] <= 59 { //对比作业后,发现要与第2个if平齐
				if FatRate >= 0.05 && FatRate < 0.12 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:偏瘦\n", Names, bmis, FatRates)
				} else if FatRate >= 0.12 && FatRate > 0.18 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:您标准健康型.\n", Names, bmis, FatRates)
				} else if FatRate >= 0.18 && FatRate > 0.23 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:您标准临界线.\n", Names, bmis, FatRates)
				} else if FatRate >= 0.23 && FatRate > 0.28 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:您标准健康型.\n", Names, bmis, FatRates)
				} else if FatRate >= 0.28 && FatRate > 0.45 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:您偏胖\n", Names, bmis, FatRates)
				} //这个代码块少了左括号,读取逻辑有变化
			} else if Ages[i] >= 60 {
				if FatRate > 0.05 && FatRate < 0.13 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:偏瘦\n", Names, bmis, FatRates)
				} else if FatRate >= 0.13 && FatRate < 0.19 {
					fmt.Printf("%s体脂率为%.2f 评估:您标准健康型.\n", Names, bmis, FatRates)
				} else if FatRate >= 0.19 && FatRate < 0.24 {
					fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:标准型警戒线.\n", Names, bmis, FatRates)
				} else if FatRate >= 0.24 && FatRate < 0.29 {
					fmt.Printf("%s体脂率为%.2f 评估:您轻度肥胖.\n", Names, bmis, FatRates)
				} else if FatRate >= 0.29 && FatRate < 0.45 {
					fmt.Printf("%s体脂率为%.2f 评估:重度肥胖\n", Names, bmis, FatRates)
				}
				//fmt.Printf(FatRates)
				//对比作业,少了左括号.少了括号,代码逻辑全部变了.所以if语句的读取顺序和条件没有和预想一致.
				//数据范围内的没有完全覆盖,所有数据无法带入if语句.
			} //对比作业,少了左括号.少了括号
		} else {
			if Ages[i] >= 18 && Ages[i] < 39 {
				if FatRate >= 0.05 && FatRate < 0.22 {
					fmt.Printf("%s体脂率为%.2f 评估:偏瘦\n", Names, bmis, FatRates)
				} else if FatRate >= 0.22 && FatRate < 0.29 {
					fmt.Printf("%s体脂率为%.2f 评估:标准健康型\n", Names, bmis, FatRates)
				} else if FatRate >= 0.29 && FatRate < 0.36 {
					fmt.Printf("%s体脂率为%.2f 评估:标准警戒线\n", Names, bmis, FatRates)
				} else if FatRate >= 0.36 && FatRate < 0.41 {
					fmt.Printf("%s体脂率为%.2f 评估:轻度肥胖\n", Names, bmis, FatRates)
				} else if FatRate >= 0.41 && FatRate < 0.45 {
					fmt.Printf("%s体脂率为%.2f 评估:重度肥胖\n", Names, bmis, FatRates)

				}
			} else if Ages[i] >= 39 && Ages[i] < 59 {
				if FatRate >= 0.05 && FatRate < 0.21 {
					fmt.Printf("%s体脂率为%.2f 评估:偏瘦\n", Names, bmis, FatRates)
				} else if FatRate >= 0.21 && FatRate < 0.28 {
					fmt.Printf("%s体脂率为%.2f 评估:标准健康型\n", Names, bmis, FatRates)
				} else if FatRate >= 0.29 && FatRate < 0.35 {
					fmt.Printf("%s体脂率为%.2f 评估:标准警戒线\n", Names, bmis, FatRates)
				} else if FatRate >= 0.35 && FatRate < 0.40 {
					fmt.Printf("%s体脂率为%.2f 评估:轻度肥胖\n", Names, bmis, FatRates)
				} else if FatRate >= 0.41 && FatRate < 0.45 {
					fmt.Printf("%s体脂率为%.2f 评估:重度肥胖\n", Names, bmis, FatRates)

				} else if Ages[i] >= 60 {
					if FatRate > 0.05 && FatRate < 0.13 {
						fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:偏瘦\n", Names, bmis, FatRates)
					} else if FatRate >= 0.13 && FatRate < 0.19 {
						fmt.Printf("%s体脂率为%.2f 评估:您标准健康型.\n", Names, bmis, FatRates)
					} else if FatRate >= 0.19 && FatRate < 0.24 {
						fmt.Printf("%sbmis指数为:%.2f,体脂率为%.2f,评估建议:标准型警戒线.\n", Names, bmis, FatRates)
					} else if FatRate >= 0.24 && FatRate < 0.29 {
						fmt.Printf("%s体脂率为%.2f 评估:您轻度肥胖.\n", Names, bmis, FatRates)
					} else if FatRate >= 0.29 && FatRate < 0.45 {
						fmt.Printf("%s体脂率为%.2f 评估:重度肥胖\n", Names, bmis, FatRates)
					}

				}

			}
		}
	}

	for i := 0; i < 3; i++ {
		totalFatRates += FatRates[i]
	}
}
