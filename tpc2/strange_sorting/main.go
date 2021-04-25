package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var testCase = 0
	var err error
	if scanner.Scan() {
		testCase, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
	}
	//
	//inputReader := bufio.NewReader(os.Stdin)
	//input, err := inputReader.ReadString('\n')
	//if err != nil {
	//	fmt.Println(err)
	//}
	//testCase, err := strconv.Atoi(input[:len(input)-1])
	//if err != nil {
	//	fmt.Println(err)
	//}

	for i := 0; i < testCase; i++ {
		//numStr, err := inputReader.ReadString('\n')
		//if err != nil {
		//	fmt.Println("222", err)
		//}
		nums := 0
		if scanner.Scan() {
			nums, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("111", err)
			}
		}

		//scanner := bufio.NewReader(os.Stdin)
		//bufio.ScanWords()

		var todo []string
		for j := 0; j < nums; j++ {
			if scanner.Scan() {
				todo = append(todo, scanner.Text())
			}
		}
		l := len(todo)
		is := false
		if l%2 != 0 {
			fmt.Print("Yes")

		} else {
			for k, _ := range todo {
				aim := (k + 2) % l
				swap(todo, k, aim)
				//fmt.Println(todo)
				if check(todo) {
					is = true
					break
				}
			}
			if is {
				fmt.Print("Yes")
			} else {
				fmt.Print("No")
			}
		}
		if i != testCase-1 {
			fmt.Print("\n")
		}
	}

}
func swap(todo []string, i, j int) {
	todo[i], todo[j] = todo[j], todo[i]
}

func check(arr []string) bool {
	if len(arr) <= 1 {
		return true
	}
	pre := arr[0]
	for k, v := range arr {
		if k == 0 {
			continue
		}
		cur, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		preNum, err := strconv.Atoi(pre)
		if err != nil {
			log.Fatal(err)
		}
		if cur < preNum {
			return false
		}
		pre = v
	}
	return true
}
