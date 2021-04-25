package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	//fmt.Println("Please enter some input: ")
	//
	//word, _ := inputReader.ReadSlice(' ')
	//fmt.Println(string(word))
	//var a,b int
	//fmt.Scanf("%d %d",&a,&b)
	//fmt.Println(a, "===",b)

	input, err := inputReader.ReadString('\n')
	if err == nil {
		//fmt.Printf("The input was: %s\n", input)
	}
	testCase, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	//var res []string
	for i := 0; i < testCase; i++ {
		_, _ = inputReader.ReadString('\n')
		//fmt.Println(caseLineNum)

		caseStr, err := inputReader.ReadString('\n')
		if err != nil {
			//fmt.Println(err)
			//return
		}
		if i == testCase-1 {
			caseStr = caseStr + "\n"
		}
		a, b := 0, 0
		isPrint := false
		cases := strings.Split(caseStr[:len(caseStr)-1], " ")
		for _, v := range cases {
			vInt, _ := strconv.Atoi(v)
			if vInt == 0 {
				fmt.Print("Yes")
				isPrint = true
				break
			}
			if vInt > 0 {
				if a > 0 {
					continue
				}
				a++
			}

			if vInt < 0 {
				if b > 0 {
					continue
				}
				b++
			}
			//fmt.Println(a,"====",b)

		}
		if !isPrint {
			if a > 0 && b > 0 {
				fmt.Print("Yes")
				//res = append(res, "Yes")
				// break
			} else {
				fmt.Print("No")
				//res = append(res,"No")
				// break
			}
		}
		if i != testCase-1 {
			fmt.Print("\n")
		}
	}

	//for _,v := range res {
	//	fmt.Println(v)
	//}

}
