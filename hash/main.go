package main

import (
	"fmt"

	"stathat.com/c/consistent"
)

func main() {
	cons := consistent.New()
	cons.Add("cacheA")
	cons.Add("cacheB")
	cons.Add("cacheC")
	server, err := cons.Get("user_89138238")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("server:", server)

	server, _ = cons.Get("user_891382383333")
	fmt.Println("server:", server)
	server, _ = cons.Get("891382383333")
	fmt.Println("server:", server)
}
