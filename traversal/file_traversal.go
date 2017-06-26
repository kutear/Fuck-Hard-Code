package traversal

import (
	"container/list"
	"fmt"
	"github.com/kutear/fuck-hard-code/parse"
	"github.com/kutear/fuck-hard-code/parse/dimens"
	jsonParse "github.com/kutear/fuck-hard-code/parse/json"
	"github.com/kutear/fuck-hard-code/parse/stringv"
	"github.com/kutear/fuck-hard-code/parse/xml"
	"github.com/kutear/fuck-hard-code/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var parseLists *list.List
var outPath string

func init() {
	parseLists = list.New()
}

func PreTraversal(json, pixel, string string, scale float64, out string) {
	strMap, dimMap := jsonParse.ParseJson(json)
	outPath = out
	utils.CreatePath(outPath)
	//DIM
	dimParse := dimens.NewDimensParse()
	dimParse.ParseDefault(pixel)
	dimParse.InitExtra(scale)
	dimParse.DealAttrs(dimMap)
	//STR
	strParse := stringv.NewStringParse()
	strParse.ParseDefault(string)
	strParse.DealAttrs(strMap)

	parseLists.PushBack(dimParse)
	parseLists.PushBack(strParse)
}

//遍历文件
func TraversalFile(layoutPath string) {
	filepath.Walk(layoutPath, walker)
	fmt.Println("遍历完成")
	saveDimOrStr()
}

func saveDimOrStr() {
	for e := parseLists.Front(); e != nil; e = e.Next() {
		p := e.Value.(parse.Parse)
		path := utils.CreatePath(outPath + string(os.PathSeparator) + "out")
		p.Save(path)
	}
}

func walker(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatalf("when traversal file : %s", err.Error())
	}
	if !info.IsDir() {
		dealFile(path, info.Name())
	}
	return nil
}

func dealFile(file string, simple string) {
	root := xml.ParseFile(file)
	dealElement(root.Element)
	filePath := outPath + string(os.PathSeparator) + simple
	ioutil.WriteFile(filePath, []byte(root.ToXML()), 0755)
}

func dealElement(element *xml.Element) {
	//属性
	for _, attr := range element.Attrs {
		parseAttrs(attr)
	}
	//子节点
	for _, child := range element.Child {
		dealElement(child)
	}
}

func parseAttrs(attr *xml.Attr) {
	for e := parseLists.Front(); e != nil; e = e.Next() {
		p := e.Value.(parse.Parse)
		if p.Match(attr.Name.Space, attr.Name.Local) {
			p.Modify(attr)
		}
	}
}
