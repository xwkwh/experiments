package main

import (
	_ "experiments/glow/os/util"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	//isTerminal := terminal.IsTerminal(int(os.Stdout.Fd()))
	//log.Println(isTerminal)
	//
	//st, err := os.Stat("main.go")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//fmt.Println(st.IsDir())
	//fmt.Println(st.Name())
	//fmt.Println(st.Mode())
	//fmt.Println(st.Size())
	//fmt.Println(st.ModTime())
	//fmt.Println(st.Sys())
	//
	//u, _ := filepath.Abs("main.go")
	//log.Println(u)
	//uu, _ := url.ParseRequestURI(u)
	//log.Printf("%+v", uu)
	//log.Printf("%+v", uu.String())
	//log.Printf("%+v", uu.Scheme)
	//log.Printf("%+v", uu.Opaque)
	//log.Printf("%+v", uu.User)
	//log.Printf("%+v", uu.Host)
	//log.Printf("%+v", uu.RawPath)
	//log.Printf("%+v", uu.Path)
	//
	//color()
	//
	//fmt.Println(os.Getenv("PAGER"))
}

func color() {
	colorFGBG := os.Getenv("COLORFGBG")
	if strings.Contains(colorFGBG, ";") {
		c := strings.Split(colorFGBG, ";")
		i, err := strconv.Atoi(c[1])
		if err == nil {
			log.Println(i)
		}
	}
	log.Println(colorFGBG)

}
