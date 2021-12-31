package calc

import "testing"

func Test1(t *testing.T) {

	fatRate, _ := calcFatRate(24.0, 30, "男")
	if fatRate != 0.192 {
		t.Fatalf("预期bmi是24，但得到的是%f", fatRate)
	}
	return

}
func Test2(t *testing.T) {

	_, err := calcFatRate(24.0, 30, "人妖")
	if err == nil {
		t.Fatalf("出现意外错误")
	}
	return

}
