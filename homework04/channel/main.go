package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Store struct {
	init  sync.Once
	store chan int
	Max   int
}

func (s *Store) instrument() {
	s.init.Do(func() {
		s.store = make(chan int, s.Max)
	})
}

type Producer struct {
	Name    string
	Fatrate uint64
}

func (r Producer) Produce(s *Store) {
	fmt.Println(r.Name, "体脂率是", r.Fatrate)
	s.store <- rand.Int()
}

type Consumer struct {
	Rank func()
}

func (Consumer) Consume(s *Store) {
	fmt.Println(Rank)
}
func Rank() {
	sort.Sort(r.Fatrate)}
	func main() {
		s := &Store{
			Max: 1000,
		}
		visitor := []Producer{}
		s.instrument()
		pCount, cCount := 1000, 1000
		for i := 0; i < pCount; i++ {
			go func() {
				for {
					rand.Seed(time.Now().UnixNano())
					visitor = append(visitor, Producer{
						Name:    fmt.Sprintf("%d", i),
						Fatrate: rand.Uint64() % 4000})
					time.Sleep(100 * time.Millisecond)
					Producer{}.Produce(s)
				}
				for range visitor {
				}
			}()
		}
		for i := 0; i < cCount; i++ {
			go func() {
				for {
					time.Sleep(100 * time.Millisecond)
					Consumer{}.Consume(s)
				}
			}()
		}
		time.Sleep(1 * time.Second)
	}
