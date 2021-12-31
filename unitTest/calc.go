package calc

import (
	"fmt"
)

func calcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {

	var sexWeight int
	if sex == "男" {
		sexWeight = 1
	} else if sex == "女" {
		sexWeight = 0
	} else {
		sexWeight = 3
	}

	fatRate, err01 := Fatrate(bmi, age, sexWeight)

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
