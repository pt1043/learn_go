package main

import "testing"

func Test_1(t *testing.T) {
	suggestion := GetSuggestion("男", 78, 20, 1.8)
	if suggestion != "目前是：偏胖。吃完饭多散散步，消化消化" {
		t.Fatalf("计算出现错误")
	}
}
