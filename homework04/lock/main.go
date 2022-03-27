package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type producer struct {
	Name    string
	Fatrate uint64
}

func (r producer) product(startPointWg, wg *sync.WaitGroup) {
	defer wg.Done()
	startPointWg.Wait()

	fmt.Println(r.Name, "体脂率是", r.Fatrate)

}

func main() {

	visitorCount := 1000
	visitor := []producer{}

	wg := sync.WaitGroup{}
	wg.Add(visitorCount)

	startPointWg := sync.WaitGroup{}
	startPointWg.Add(1)

	for i := 0; i < visitorCount; i++ {
		rand.Seed(time.Now().UnixNano())
		visitor = append(visitor, producer{
			Name:    fmt.Sprintf("%d", i),
			Fatrate: rand.Uint64() % 4000,
		})
	}

	for _, visitorItem := range visitor {
		go visitorItem.product(&startPointWg, &wg)
	}

	startPointWg.Done()

	wg.Wait()

}
