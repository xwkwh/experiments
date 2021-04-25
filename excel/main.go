package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type T []Feat

type Feat struct {
	Path     string `yaml:"path"`
	SlotID   int    `yaml:"slot_id"`
	Type     string `yaml:"type"`
	FeatType int    `yaml:"feat_type"`
}

var featMap = map[string]int{
	"numeric":         0,
	"identity":        1,
	"categorical":     2,
	"listcategorical": 3,
}

func main() {
	fmt.Println("main.....................")
	f, err := excelize.OpenFile("feat.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("Sheet1")
	var t T
	for _, row := range rows {
		feat := Feat{
			Path:     row[2],
			FeatType: featMap[strings.ToLower(row[3][:strings.Index(row[3], "(")])],
		}
		slotID, _ := strconv.Atoi(row[0])
		feat.SlotID = slotID
		if row[2][:4] == "user" {
			feat.Type = "user"
		} else {
			feat.Type = "item"
		}
		t = append(t, feat)
	}
	res, err := yaml.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
	ioutil.WriteFile("feat.yaml", res, os.ModePerm)
	//pack.Pack()

}
