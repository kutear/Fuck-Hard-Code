package stringv

import (
	e_xml "encoding/xml"
	"fmt"
	"github.com/kutear/fuck-hard-code/parse"
	"github.com/kutear/fuck-hard-code/parse/json"
	"github.com/kutear/fuck-hard-code/parse/xml"
	"github.com/kutear/fuck-hard-code/utils"
	"log"
	"os"
	"strings"
)

type StringParse struct {
	attrs  []json.ConfigItem   //可以解析的属性
	values map[string](string) //存在的值
	index  int
}

func NewStringParse() *StringParse {
	return &StringParse{
		values: make(map[string](string), 0),
	}
}

func (this *StringParse) ParseDefault(path string) {
	parse.ParseValue(path, &parse.XmlString{}, this.values)
}

func (this *StringParse) InitExtra(extra interface{}) {
}

func (this *StringParse) DealAttrs(attr interface{}) {
	if data, ok := attr.([]json.ConfigItem); ok {
		this.attrs = data
	}
}

func (this *StringParse) Match(spaceName string, attrsName string) bool {
	for _, item := range this.attrs {
		if item.NameSpace == spaceName && utils.Contains(item.Items, attrsName) {
			return true
		}
	}
	return false
}

func (this *StringParse) Modify(attr *xml.Attr) {
	value := attr.Value
	if isHardCodeStr(value) {
		if dim, ok := this.values[value]; ok {
			attr.Value = fmt.Sprintf("@string/%s", dim)
		} else {
			//构建这个值
			buildKey := this.buildStringItem()
			this.values[value] = buildKey
			attr.Value = fmt.Sprintf("@string/%s", buildKey)
		}
	}
}

func (this *StringParse) buildStringItem() string {
	for true {
		ret := fmt.Sprintf("str_%d", this.index)
		this.index++
		if _, ok := this.values[ret]; !ok {
			return ret
		}
	}
	return ""
}

func (this *StringParse) Save(path string) {
	ret := utils.SortMap(this.values)
	strs := parse.XmlString{
		Tags: ret,
	}
	byteData, err := e_xml.MarshalIndent(strs, "  ", "    ")
	if err != nil {
		log.Fatalf("输出strings.xml失败 %s", err.Error())
	}
	utils.Write(byteData, path+string(os.PathSeparator)+"strings.xml")
}

func isHardCodeStr(string string) bool {
	return !strings.HasPrefix(string, "@string/")
}
