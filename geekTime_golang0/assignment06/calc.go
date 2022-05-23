package main

import (
	"fmt"
	"log"
	"math"
	"sync"

	"learn.go/pkg/apis"
)

var _ ServeInterface = &FatRateItem{}

type RandItem struct {
	Name    string
	Sex     string
	fatRate float64
}

type FatRateItem struct {
	items     []RandItem
	itemsLock sync.Mutex
}

func (r *FatRateItem) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	bmi, err := BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Println("计算BMI失败：", err)
		return err
	}
	fr := CalcFatRate(bmi, int(pi.Age), pi.Sex)
	r.inputRecord(pi.Name, pi.Sex, fr)
	return nil
}

func (r *FatRateItem) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	bmi, err := BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Println("计算BMI失败：", err)
		return nil, err
	}
	fr := CalcFatRate(bmi, int(pi.Age), pi.Sex)
	r.inputRecord(pi.Name, pi.Sex, fr)
	return &apis.PersonalInformationFatRate{
		Name:    pi.Name,
		FatRate: float32(fr),
	}, nil
}

func NewFatRateRank() *FatRateItem {
	return &FatRateItem{
		items: make([]RandItem, 0, 100),
	}
}

func (r *FatRateItem) inputRecord(name, sex string, fatRate ...float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.fatRate >= minFatRate {
				item.fatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, RandItem{
			Name:    name,
			Sex:     sex,
			fatRate: minFatRate,
		})
	}
}

func BMI(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("weight cannot be negative")
		return
	}
	if heightM < 0 {
		err = fmt.Errorf("height cannot be negative")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("height cannot be 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}
func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	fatRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
