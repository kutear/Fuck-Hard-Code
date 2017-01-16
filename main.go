package main

import (
	"encoding/xml"
	"fmt"
	"github.com/kutear/Fuck-Hard-Code/parser"
	"io/ioutil"
)

func main() {
	xmlFile, err := ioutil.ReadFile("/home/kutear/Desktop/test2.xml")
	if err != nil {
		fmt.Println(err.Error())
		panic("读取文件失败")
	}
	fmt.Println(string(xmlFile))
	s := &parser.DimensionComponent{}
	xml.Unmarshal([]byte(xmlFile), &s)
	fmt.Println(s)
}
