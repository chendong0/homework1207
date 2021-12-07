package main

import (
	"fmt"
)

//全局变量声明
var (
	avearageFat float64 = 0.0 //平均体脂
	peopleCount int     = 0   //录入人数计数器
	isContinue  bool          //循环采集是否继续

	Weight    float64 //体重(KG)
	Tall      float64 //身高(m)
	Age       int     //年龄
	SexWeight int     //性别权重取值
	Sex       string  //性别
	fatRate   float64 //体脂率
	bmi       float64 //bmi值
	Name      string  //姓名

)

func main() {

	fmt.Println("hello69") //启动信息

	for isContinue = true; isContinue; {
		getInfor()
		peopleCount++ //++加号没有空格隔开.
		var bmi = (Weight / (Tall * Tall))
		var (
			fatRate float64 = (1.2*bmi + 0.23*float64(Age) - 5.4 - 10.8*float64(SexWeight)) / 100
		)
		fmt.Printf("%s的bmi指数为%.2f \n", Name, bmi)
		fmt.Printf("%s的体脂率为%.2f \n", Name, fatRate)

		calculator()
		fmt.Printf("+----------+\n累计录入%d人记录,平均体脂为%.2f \n"+"+----------+\n",
			peopleCount, (avearageFat*float64(peopleCount-1)+fatRate)/float64(peopleCount))
		//exitbmi()

		var whetherContinue string
		fmt.Print("\n是否录入下一个(y/n)?")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}

func getInfor() {
	fmt.Print("\n姓名:")
	fmt.Scanln(&Name)

	fmt.Print("体重(KG):")
	fmt.Scanln(&Weight)

	fmt.Print("身高(米):")
	fmt.Scanln(&Tall)

	fmt.Print("年龄:")
	fmt.Scanln(&Age)

	fmt.Println("性别(female/male):")
	fmt.Scanln(&Sex)

	if Sex == "male" { //==要区分理解=和==的区别和用法.
		SexWeight = 0
	} else {
		SexWeight = 1
	}
}

func calculator() {
	fmt.Println("+++++---+++") //打印测试

	if Age < 18 {
		fmt.Println("不评估18岁以下,因为波动太大,效果没有意义")
	} //else if嵌套使用,熟料度,和练习相互嵌套
	//要把共有条件先放在前面
	//func calculator() {
	//	fmt.Println("+++++++++++-----------+++++++++++")
	//		}

	if Sex == "male" {
		//编写男性体脂率与体脂状态

		if Age >= 18 && Age <= 39 {
			if fatRate < 0.05 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:偏瘦,加强营养多锻炼\n", Name, bmi, fatRate)
				//%.2f 评估双引号里面的空格也会导致读取数据错误
			} else if fatRate >= 0.05 && fatRate < 0.1 { //5%报错,只能用0.15
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:您的状态偏瘦\n", Name, bmi, fatRate)
			} else if fatRate >= 0.1 && fatRate < 0.17 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:您目前标准,请保持\n", Name, bmi, fatRate)
			} else if fatRate >= 0.17 && fatRate < 0.22 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:您达到标准警戒线\n", Name, bmi, fatRate)

			} else if fatRate >= 0.22 && fatRate < 0.27 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:您轻度肥胖\n", Name, bmi, fatRate)
			} else if fatRate >= 0.27 && fatRate < 0.45 {
				fmt.Printf("BMI指数为:%.2f,体脂率为%.2f,评估建议:您重度肥胖\n", bmi, fatRate)
			} //对比作业,多加了左括号

		} else if Age > 39 && Age <= 59 { //对比作业后,发现要与第2个if平齐
			if fatRate >= 0.05 && fatRate < 0.12 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:偏瘦\n", Name, bmi, fatRate)
			} else if fatRate >= 0.12 && fatRate > 0.18 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:您标准健康型.\n", Name, bmi, fatRate)
			} else if fatRate >= 0.18 && fatRate > 0.23 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:您标准临界线.\n", Name, bmi, fatRate)
			} else if fatRate >= 0.23 && fatRate > 0.28 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:您标准健康型.\n", Name, bmi, fatRate)
			} else if fatRate >= 0.28 && fatRate > 0.45 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:您偏胖\n", Name, bmi, fatRate)
			} //这个代码块少了左括号,读取逻辑有变化
		} else if Age >= 60 {
			if fatRate > 0.05 && fatRate < 0.13 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:偏瘦\n", Name, bmi, fatRate)
			} else if fatRate >= 0.13 && fatRate < 0.19 {
				fmt.Printf("%s体脂率为%.2f 评估:您标准健康型.\n", Name, bmi, fatRate)
			} else if fatRate >= 0.19 && fatRate < 0.24 {
				fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:标准型警戒线.\n", Name, bmi, fatRate)
			} else if fatRate >= 0.24 && fatRate < 0.29 {
				fmt.Printf("%s体脂率为%.2f 评估:您轻度肥胖.\n", Name, bmi, fatRate)
			} else if fatRate >= 0.29 && fatRate < 0.45 {
				fmt.Printf("%s体脂率为%.2f 评估:重度肥胖\n", Name, bmi, fatRate)
			}
			//fmt.Printf(fatRate)
			//对比作业,少了左括号.少了括号,代码逻辑全部变了.所以if语句的读取顺序和条件没有和预想一致.
			//数据范围内的没有完全覆盖,所有数据无法带入if语句.
		} //对比作业,少了左括号.少了括号
	} else {
		if Age >= 18 && Age < 39 {
			if fatRate >= 0.05 && fatRate < 0.22 {
				fmt.Printf("%s体脂率为%.2f 评估:偏瘦\n", Name, bmi, fatRate)
			} else if fatRate >= 0.22 && fatRate < 0.29 {
				fmt.Printf("%s体脂率为%.2f 评估:标准健康型\n", Name, bmi, fatRate)
			} else if fatRate >= 0.29 && fatRate < 0.36 {
				fmt.Printf("%s体脂率为%.2f 评估:标准警戒线\n", Name, bmi, fatRate)
			} else if fatRate >= 0.36 && fatRate < 0.41 {
				fmt.Printf("%s体脂率为%.2f 评估:轻度肥胖\n", Name, bmi, fatRate)
			} else if fatRate >= 0.41 && fatRate < 0.45 {
				fmt.Printf("%s体脂率为%.2f 评估:重度肥胖\n", Name, bmi, fatRate)

			}
		} else if Age >= 39 && Age < 59 {
			if fatRate >= 0.05 && fatRate < 0.21 {
				fmt.Printf("%s体脂率为%.2f 评估:偏瘦\n", Name, bmi, fatRate)
			} else if fatRate >= 0.21 && fatRate < 0.28 {
				fmt.Printf("%s体脂率为%.2f 评估:标准健康型\n", Name, bmi, fatRate)
			} else if fatRate >= 0.29 && fatRate < 0.35 {
				fmt.Printf("%s体脂率为%.2f 评估:标准警戒线\n", Name, bmi, fatRate)
			} else if fatRate >= 0.35 && fatRate < 0.40 {
				fmt.Printf("%s体脂率为%.2f 评估:轻度肥胖\n", Name, bmi, fatRate)
			} else if fatRate >= 0.41 && fatRate < 0.45 {
				fmt.Printf("%s体脂率为%.2f 评估:重度肥胖\n", Name, bmi, fatRate)

			} else if Age >= 60 {
				if fatRate > 0.05 && fatRate < 0.13 {
					fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:偏瘦\n", Name, bmi, fatRate)
				} else if fatRate >= 0.13 && fatRate < 0.19 {
					fmt.Printf("%s体脂率为%.2f 评估:您标准健康型.\n", Name, bmi, fatRate)
				} else if fatRate >= 0.19 && fatRate < 0.24 {
					fmt.Printf("%sBMI指数为:%.2f,体脂率为%.2f,评估建议:标准型警戒线.\n", Name, bmi, fatRate)
				} else if fatRate >= 0.24 && fatRate < 0.29 {
					fmt.Printf("%s体脂率为%.2f 评估:您轻度肥胖.\n", Name, bmi, fatRate)
				} else if fatRate >= 0.29 && fatRate < 0.45 {
					fmt.Printf("%s体脂率为%.2f 评估:重度肥胖\n", Name, bmi, fatRate)
				}

			}

		}
	}
}

//func exitbmi() {
//
//	}
//}
