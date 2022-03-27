package main

import (
	"fmt"
	"time"
)

func main() {

	getmessage()

}

func inputfloor() (int, int, int) {
	var obfloor [5]int
	fmt.Println("输入楼层: ")

	fmt.Scanln(&obfloor[0], &obfloor[1], &obfloor[2])

	return obfloor[0], obfloor[1], obfloor[2]
}
func elevatorfloor() int {

	var nowfloor int

	fmt.Println("扫描当前电梯位置")

	fmt.Scanln(&nowfloor)

	return nowfloor
}
func getmessage() {

	var button1, button2, button3 = inputfloor()

	var checkfloor = elevatorfloor()

	if button1 > checkfloor && button1 > button2 && button1 > button3 && button2 > button3 && button3 < checkfloor {
		fmt.Println("电梯上行")
		time.Sleep(100)
		fmt.Println("电梯到达", button1, "楼")
		time.Sleep(100)
		fmt.Println("电梯下行")
		time.Sleep(100)
		fmt.Println("电梯到达", button3, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button2, "楼")
	}
	if button1 > checkfloor && button1 < button2 && button1 > button3 && button2 > button3 && button3 < checkfloor {
		fmt.Println("电梯上行")
		time.Sleep(100)
		fmt.Println("电梯到达", button1, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button2, "楼")
		time.Sleep(100)
		fmt.Println("电梯下行")
		time.Sleep(100)
		fmt.Println("电梯到达", button3, "楼")

	}
	if button1 > checkfloor && button1 < button2 && button1 < button3 && button2 > button3 {
		fmt.Println("电梯上行")
		time.Sleep(100)
		fmt.Println("电梯到达", button1, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button3, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button2, "楼")

	}
	if button1 < checkfloor && button1 > button2 && button1 > button3 && button2 < button3 {
		fmt.Println("电梯下行")
		time.Sleep(100)
		fmt.Println("电梯到达", button1, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button3, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button2, "楼")
	}
	if button1 < checkfloor && button1 > button2 && button1 > button3 && button2 > button3 {
		fmt.Println("电梯下行")
		time.Sleep(100)
		fmt.Println("电梯到达", button1, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button2, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button3, "楼")
	}
	if button1 < checkfloor && button1 < button2 && button1 < button3 && button2 > button3 {
		fmt.Println("电梯下行")
		time.Sleep(100)
		fmt.Println("电梯到达", button2, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button3, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button1, "楼")
	}
	if button1 < checkfloor && button1 < button2 && button1 > button3 && button2 > button3 {
		fmt.Println("电梯下行")
		time.Sleep(100)
		fmt.Println("电梯到达", button2, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button1, "楼")
		time.Sleep(100)
		fmt.Println("电梯到达", button3, "楼")
	}
}
