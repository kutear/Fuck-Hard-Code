package main

import (
	"./utils"
	_ "./utils"
	"fmt"
	"io/ioutil"
)

func main() {
	xmlFile, err := ioutil.ReadFile("/home/kutear/Desktop/test2.xml")
	if err != nil {
		fmt.Println(err.Error())
		panic("读取文件失败")
	}
	element, err := utils.LoadByXml(string(xmlFile))
	if err != nil {

	}
	utils.DealElement(element, "test.xml")
}
