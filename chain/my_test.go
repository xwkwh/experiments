package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_My(t *testing.T) {
	a := time.Now()
	time.Sleep(time.Duration(500))
	fmt.Println(time.Since(a))
}
