package gobmi

import "testing"

func TestBMI1(t *testing.T) {

	bmi, err := BMI(78.0, 1.8)
	var Bmiend = int(bmi)
	if Bmiend != 24 {
		t.Fatalf("预期bmi是24，但得到的是%f", bmi)
	}
	return
	if err != nil {
		t.Fatalf("出现意外错误")

	}

}
func TestBMI2(t *testing.T) {

	_, err := BMI(78.0, 0)
	if err == nil {
		t.Fatalf("出现意外错误")

	}

}
func TestBMI3(t *testing.T) {

	_, err := BMI(0, 1.8)
	if err == nil {
		t.Fatalf("出现意外错误")

	}

}
