package main

import (
	"fmt"
)

func GetSuggestion(sex string, weightKG float64, age int, heightM float64) (suggestion string) {

	fatRate, err := calcFatRate(weightKG, heightM, age, sex)

	if err != nil {
		return
	}

	answer1 := getHealthinessSuggestionsForMale(age, fatRate)

	answer := getHealthinessSuggestionsForFeMale(age, fatRate)
	if sex == "男" {
		suggestion = answer1
	} else {
		suggestion = answer
	}
	fmt.Println(suggestion)
	return
}

func calcFatRate(weightKG float64, heightM float64, age int, sex string) (fatRate float64, err error) {

	var sexWeight int
	if sex == "男" {
		sexWeight = 1
	} else if sex == "女" {
		sexWeight = 0
	} else {
		sexWeight = 3
	}
	bmi, err := BMI(weightKG, heightM)

	fatRate, err01 := Fatrate(bmi, age, sexWeight)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if err01 != nil {
		fmt.Println(err01)
		return 0, err01
	}
	fmt.Println("体脂率是：", fatRate)

	return fatRate, nil

}
func Fatrate(bmi float64, age int, sexWeight int) (fatrate float64, err01 error) {
	if age < 0 {
		err01 = fmt.Errorf("age cannot be negative")
		return 0, err01
	}
	if age > 150 {
		err01 = fmt.Errorf("age cannot be negative")
		return 0, err01

	}
	if sexWeight == 3 {
		err01 = fmt.Errorf("sex cannot be negative")
		return 0, err01
	}

	fatrate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100

	return fatrate, nil
}
func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
func BMI(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("weight cannot be negative")
		return
	}
	if heightM < 0 {
		err = fmt.Errorf("height cannot be negative")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("height cannot be 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}

func getHealthinessSuggestionsForMale(age int, fatRate float64) (answer1 string) {
	//  编写男性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			answer1 = ("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			answer1 = ("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			answer1 = ("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			answer1 = ("目前是：肥胖。少吃点，多运动运动")
		} else {
			answer1 = ("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.11 {
			answer1 = ("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.11 && fatRate <= 0.17 {
			answer1 = ("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.17 && fatRate <= 0.22 {
			answer1 = ("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.22 && fatRate <= 0.27 {
			answer1 = ("目前是：肥胖。少吃点，多运动运动")
		} else {
			answer1 = ("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 60 {
		if fatRate <= 0.13 {
			answer1 = ("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.13 && fatRate <= 0.19 {
			answer1 = ("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.19 && fatRate <= 0.24 {
			answer1 = ("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.24 && fatRate <= 0.29 {
			answer1 = ("目前是：肥胖。少吃点，多运动运动")
		} else {
			answer1 = ("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else {
		answer1 = ("我们不看未成年人的体脂率，变化太大，无法评判")
	}
	return answer1
}
func getHealthinessSuggestionsForFeMale(age int, fatRate float64) (answer string) {
	//  编写女性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		if fatRate <= 0.2 {
			answer = ("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.2 && fatRate <= 0.27 {
			answer = ("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.27 && fatRate <= 0.34 {
			answer = ("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.34 && fatRate <= 0.39 {
			answer = ("目前是：肥胖。少吃点，多运动运动")
		} else {
			answer = ("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.21 {
			answer = ("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.21 && fatRate <= 0.28 {
			answer = ("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.28 && fatRate <= 0.35 {
			answer = ("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.35 && fatRate <= 0.40 {
			answer = ("目前是：肥胖。少吃点，多运动运动")
		} else {
			answer = ("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 60 {
		if fatRate <= 0.22 {
			answer = ("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.22 && fatRate <= 0.29 {
			answer = ("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.29 && fatRate <= 0.36 {
			answer = ("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.36 && fatRate <= 0.41 {
			answer = ("目前是：肥胖。少吃点，多运动运动")
		} else {
			answer = ("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else {
		answer = ("我们不看未成年人的体脂率，变化太大，无法评判")
	}
	return answer
}
