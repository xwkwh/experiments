package main

import (
	"log"
	"os"
	"reflect"
)

func main() {
	Stat("/Users/tobyguo/gopath/src/experiments/swag/os/main.go")
	//fmt.Print(111)
	Stat("./main.go")
}

func Stat(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf("dir: %s is not exist\n", dir)
		log.Println(err)
		log.Println(reflect.TypeOf(err))

		er := underlyingError(err)
		log.Println(reflect.TypeOf(er))
		log.Println(er)

	}
}

func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	}
	return err
}
