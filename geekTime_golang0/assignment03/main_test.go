package main

import "testing"

func TestCase1(t *testing.T) {

	inputfloor(nil)
	Final := getmessage(nil)
	if Final != "电梯不动" {
		t.Fatalf("出现错误")
	}

}
func TestCase2(t *testing.T) {

	inputfloor(3)
	elevatorfloor(1)
	Final := getmessage
	if Final != "电梯到达3楼" {
		t.Fatalf("出现错误")
	}

}
func TestCase3(t *testing.T) {

	inputfloor(4, 2)
	elevatorfloor(3)
	Final := getmessage
	if Final != "电梯到达2楼" {
		t.Fatalf("出现错误")
	}

}
func TestCase4(t *testing.T) {

	inputfloor(4, 5, 2)
	elevatorfloor(3)
	Final := getmessage
	if Final != "电梯到达2楼" {
		t.Fatalf("出现错误")
	}
}
