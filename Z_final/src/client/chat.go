package main

import (
	"context"
	"fmt"
	"io"
	"learn.go/pkg/apis"
	"log"
	"time"
)

var information = [3]apis.PersonalInformation

func Chatservice(lis context.Context) {
	defer fmt.Println("结束链接：", lis)
	defer lis.Close()
	for {
		buf := make([]byte, 1024)
		valid, err := lis.Read(Context)
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("WARNING：读数据失败：", err)
			continue
		}

		fmt.Println([3]apis.PersonalInformation)
	}
}
func  GetInput() *apis.PersonalInformation {
	var message string
	fmt.Scanln(&message)

	return &apis.PersonalInformation{
		Message:   message,
	}}