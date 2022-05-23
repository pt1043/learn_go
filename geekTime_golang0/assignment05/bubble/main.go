package main

import (
	"fmt"
	"log"
)

func main() {
	input := &inputFromStd{}
	calc := &Calc{}
	rk := &FatRateRank{}
	records := NewRecord("C:\\Users\\78607\\Desktop\\records.txt")

	for {
		pi := input.GetInput()
		if err := records.savePersonalInformation(pi); err != nil {
			log.Fatal("保存失败：", err)
		}
		fr, err := calc.FatRate(pi)
		if err != nil {
			log.Fatal("计算体脂率失败：", err)
		}
		rk.inputRecord(pi.Name, fr)

		rankResult, _ := rk.getRank(pi.Name)
		fmt.Println("排名结果：", rankResult)

		if cont := whetherContinue(); !cont {
			break
		}
	}
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
