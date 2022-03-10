package main

import (
	add "add"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {
	for {
		mainFatRateBody()

		if cont := whetherContinue(); !cont {
			break
		}
	}
}
func calcFatRate(weightKG float64, heightM float64, age int, sexWeight int) (fatRate float64, err error) {

	bmi, err := gobmi.BMI(weightKG, heightM)

	fatRate, err01 := add.Fatrate(bmi, age, sexWeight)
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
func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
func mainFatRateBody() {

	_, weightKG, heightM, age, sex, sexWeight := getMaterialsFromInput()

	fatRate, err := calcFatRate(weightKG, heightM, age, sexWeight)

	if err != nil {
		return
	}

	var checkHealthinessFunc func(age int, fatRate float64)
	if sex == "男" {
		checkHealthinessFunc = getHealthinessSuggestionsForMale
	} else {
		checkHealthinessFunc = getHealthinessSuggestionsForFeMale
	}
	getHealthinessSuggestions(age, fatRate, checkHealthinessFunc)
}
func getHealthinessSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	getSuggestion(age, fatRate)
}
func getMaterialsFromInput() (string, float64, float64, int, string, int) {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weightKG float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weightKG)

	var heightM float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&heightM)
	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	var sex string
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)
	var sexWeight int
	if sex == "男" {
		sexWeight = 1
	} else if sex == "女" {
		sexWeight = 0
	} else {
		sexWeight = 3
	}
	return name, weightKG, heightM, age, sex, sexWeight
}
func whetherContinue() bool {
	var whetherContinue string
	fmt.Print("是否录入下一个（y/n）？")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}
