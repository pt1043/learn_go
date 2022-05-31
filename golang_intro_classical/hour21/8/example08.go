package main

import (
	"fmt"
	"github.com/Burntsushi/toml"
	"log"
)

type Config struct {
	Name   string
	Awake  bool
	Hungry bool
}

func main() {
	c := Config{}
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}
