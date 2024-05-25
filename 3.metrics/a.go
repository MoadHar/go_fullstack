package main

import (
	"errors"
	"log"
)

type ShutdownFunc func(a int) error

func initFoo(service string) (ShutdownFunc, error) {
	return func(a int) error { return errors.New("ahaaa") }, errors.New("errror")
}

func main() {
	v, err := initFoo("something in the way")
	log.Println(v(1785))
	log.Println(err)
}
