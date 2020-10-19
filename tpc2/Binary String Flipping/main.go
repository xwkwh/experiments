package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	total := 0
	fmt.Scanln( &total)
	inputReader := bufio.NewReader(os.Stdin)

	for i:=0;i<total;i++ {
		input ,_,_ :=inputReader.ReadLine()

		l := len(string(input))

		inputMap := make(map[byte]int)
		inputStr := string(input)
		for _,v := range inputStr {
			if v =='1' {
				inputMap['1']++
			}else {
				inputMap['0']++
			}
		}
		if inputMap['1'] - inputMap['0']<=1 {
			fmt.Println(1,l)
		}else {
			for i:=len(inputStr);i>0;i++ {
				inputMap[inputStr[i-1]]--
				if inputMap['1'] - inputMap['0']<=1  {
					fmt.Println(1,i)
					return
				}
			}
		}
	}
}
