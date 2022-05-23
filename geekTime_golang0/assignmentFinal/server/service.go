package main

import (
	context2 "golang.org/x/net/context"
	"io"
	"learn.go/pkg/apis"
	"log"
	"sync"
)

var _ apis.Chatservice = &chatservice{}

type chatServer struct {
	sync.Mutex
	persons  map[string]*apis.PersonalInformation
	persons  map[string]*apis.PersonalInformation
	personCh chan *apis.PersonalInformation
}

func (r *Chatservice) regPerson(pi *apis.PersonalInformation) {
	r.Lock()
	defer r.Unlock()
	r.persons[pi.Name] = pi
	r.personCh <- pi
}

func (r *rankServer) WatchPersons(null *apis.Null, server apis.RankService_WatchPersonsServer) error {
	for pi := range r.personCh {
		if err := server.Send(pi); err != nil {
			log.Println("发送失败，结束:", err)
			return err
		}
	}
	return nil
}

func (r *rankServer) RegisterPersons(server apis.RankService_RegisterPersonsServer) error {
	pis := &apis.PersonalInformationList{}
	for {
		pi, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Printf("WARNING: 获取人员注册时失败：%v\n", err)
			return err
		}

	}
	return nil
}
