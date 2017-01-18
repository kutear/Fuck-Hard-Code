package main

import (
	_ "./utils"
	"fmt"
	"github.com/kutear/Fuck-Hard-Code/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	//
	//inPath := flag.String("input", "D:\\", "Layout根文件目录")
	////outPath := flag.String("output", "D:\\", "文件输出目录")
	//flag.Parse()
	inPath := "G:\\ll"
	outPath := "G:\\layout2"
	utils.CreatePath(outPath)
	err := filepath.Walk(inPath, func(file string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println("跳过目录" + file)
			return nil
		}
		utils.DealFile(file, info.Name(), outPath)
		return nil
	})

	ioutil.WriteFile(outPath+string(os.PathSeparator)+"dimens.xml", []byte(utils.GetDimen().ToXML()), 0644)
	ioutil.WriteFile(outPath+string(os.PathSeparator)+"strings.xml", []byte(utils.GetString().ToXML()), 0644)

	if err != nil {
		fmt.Println(err.Error())
	}

}
