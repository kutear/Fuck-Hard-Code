package main

import (
	"io/ioutil"
	"fmt"
	"encoding/xml"
	"github.com/kutear/Fuck-Hard-Code/parser"
)

//type DimensionParser struct {
//	Width  string `xml:"android:layout_width,attr"`
//	Height string `xml:"android:layout_height,attr"`
//}

func main() {
	xmlFile, err := ioutil.ReadFile("/home/kutear/Desktop/test.xml")
	if err != nil {
		fmt.Println(err.Error())
		panic("读取文件失败")
	}
	fmt.Println(string(xmlFile))
	s := &parser.DimensionParser{}
	xml.Unmarshal([]byte(xmlFile), &s)
	fmt.Println(s)
}
