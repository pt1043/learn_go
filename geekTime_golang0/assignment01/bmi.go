package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		var name string
		fmt.Print("名字：")
		fmt.Scanln(&name)

		var weight float32
		fmt.Print("体重（kg）：")
		fmt.Scanln(&weight)

		var tall float32
		fmt.Print("身高（米）：")
		fmt.Scanln(&tall)

		var bmi float32 = weight / (tall * tall)

		var age int
		fmt.Print("年龄（岁）：")
		fmt.Scanln(&age)

		var sexWeight int
		var sex string = "男"
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)

		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		var fatRate float32 = (1.2*bmi + 0.23*float32(age) - 5.4 - 10.8*float32(sexWeight)) / 100
		fmt.Println("体脂率是：", fatRate)

		if sex == "男" {
			//男性的体脂率和状态
			if age >= 18 && age <= 39 {
				if fatRate <= 0.1 {
					fmt.Println("偏瘦.")
				} else if fatRate > 0.1 && fatRate <= 0.16 {
					fmt.Println("标准.")
				} else if fatRate > 0.16 && fatRate <= 0.21 {
					fmt.Println("偏胖.")
				} else if fatRate > 0.21 && fatRate <= 0.26 {
					fmt.Println("肥胖.")
				} else {
					fmt.Println("非常肥胖，需注意生命质量.")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.11 {
					fmt.Println("偏瘦.")
				} else if fatRate > 0.11 && fatRate <= 0.17 {
					fmt.Println("标准.")
				} else if fatRate > 0.17 && fatRate <= 0.22 {
					fmt.Println("偏胖.")
				} else if fatRate > 0.22 && fatRate <= 0.27 {
					fmt.Println("肥胖.")
				} else {
					fmt.Println("非常肥胖，需注意生命质量.")
				}
			} else if age >= 60 {
				if fatRate <= 0.13 {
					fmt.Println("偏瘦.")
				} else if fatRate > 0.13 && fatRate <= 0.19 {
					fmt.Println("标准.")
				} else if fatRate > 0.19 && fatRate <= 0.24 {
					fmt.Println("偏胖.")
				} else if fatRate > 0.24 && fatRate <= 0.29 {
					fmt.Println("肥胖.")
				} else {
					fmt.Println("非常肥胖，需注意生命质量.")
				}
			} else {
				fmt.Println("是未成年人或录入严重错误的体重和身高在计算公式标准差范围内计算不准确，所以不做判断……")
			}

		} else {
			if age >= 18 && age <= 39 {
				if fatRate <= 0.2 {
					fmt.Println("偏瘦.")
				} else if fatRate > 0.2 && fatRate <= 0.27 {
					fmt.Println("标准.")
				} else if fatRate > 0.27 && fatRate <= 0.34 {
					fmt.Println("偏胖.")
				} else if fatRate > 0.34 && fatRate <= 0.39 {
					fmt.Println("肥胖.")
				} else {
					fmt.Println("非常肥胖，需注意生命质量.")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.21 {
					fmt.Println("偏瘦.")
				} else if fatRate > 0.21 && fatRate <= 0.28 {
					fmt.Println("标准.")
				} else if fatRate > 0.28 && fatRate <= 0.35 {
					fmt.Println("偏胖.")
				} else if fatRate > 0.35 && fatRate <= 0.40 {
					fmt.Println("肥胖.")
				} else {
					fmt.Println("非常肥胖，需注意生命质量.")
				}
			} else if age >= 60 {
				if fatRate <= 0.22 {
					fmt.Println("偏瘦.")
				} else if fatRate > 0.22 && fatRate <= 0.29 {
					fmt.Println("标准.")
				} else if fatRate > 0.29 && fatRate <= 0.36 {
					fmt.Println("偏胖.")
				} else if fatRate > 0.36 && fatRate <= 0.41 {
					fmt.Println("肥胖.")
				} else {
					fmt.Println("非常肥胖，需注意生命质量.")
				}
			} else {
				fmt.Println("是未成年人或录入严重错误的体重和身高在计算公式标准差范围内计算不准确，所以不做判断……")
			}
		}

		var whetherContinue string
		fmt.Print("是否录入下一个（Y/n）？")
		fmt.Scanln(&whetherContinue)

		if whetherContinue != "y" {
			break
		}

	}
}
