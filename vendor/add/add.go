package add

import "fmt"

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
