package dimens

import (
	e_xml "encoding/xml"
	"fmt"
	"github.com/kutear/fuck-hard-code/parse"
	"github.com/kutear/fuck-hard-code/parse/json"
	"github.com/kutear/fuck-hard-code/parse/xml"
	"github.com/kutear/fuck-hard-code/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

//sub class Parse
type DimensParse struct {
	attrs      []json.ConfigItem   //可以解析的属性
	values     map[string](string) //存在的值
	scaleRatio float64
}

func NewDimensParse() *DimensParse {
	return &DimensParse{
		values:     make(map[string](string), 0),
		scaleRatio: 3,
	}
}

func (this *DimensParse) ParseDefault(path string) {
	parse.ParseValue(path, &parse.XmlDimens{}, this.values)
}

func (this *DimensParse) InitExtra(extra interface{}) {
	if data, ok := extra.(float64); ok {
		this.scaleRatio = data
	}
}

func (this *DimensParse) DealAttrs(attr interface{}) {
	if data, ok := attr.([]json.ConfigItem); ok {
		this.attrs = data
	}
}

func (this *DimensParse) Match(spaceName string, attrsName string) bool {
	for _, item := range this.attrs {
		if item.NameSpace == spaceName && utils.Contains(item.Items, attrsName) {
			return true
		}
	}
	return false
}

func (this *DimensParse) Modify(attr *xml.Attr) {
	value := attr.Value
	if isHardCodeDimen(value) {
		value = strings.Replace(value, "dip", "dp", -1)
		suf := value[len(value)-2:] //px or dp and sp
		val := value[:len(value)-2]
		if suf == "px" {
			if temp, err := strconv.Atoi(val); err == nil {
				ret := float64(temp) / this.scaleRatio
				value = fmt.Sprintf("%.1f%s", ret, "dp")
				suf = value[len(value)-2:]
				val = value[:len(value)-2]
			}
		}
		if dim, ok := this.values[value]; ok {
			attr.Value = fmt.Sprintf("@dimen/%s", dim)
		} else {
			//构建这个值
			buildKey := buildDimenItem(value)
			this.values[value] = buildKey
			attr.Value = fmt.Sprintf("@dimen/%s", buildKey)
		}
	}
}

func (this *DimensParse) Save(path string) {
	ret := utils.SortMap(this.values)
	xmls := parse.XmlDimens{
		Tags: ret,
	}
	byteData, err := e_xml.MarshalIndent(xmls, "  ", "    ")
	if err != nil {
		log.Fatalf("输出dimems.xml失败 %s", err.Error())
	}
	utils.Write(byteData, path+string(os.PathSeparator)+"dimens.xml")
}

func buildDimenItem(value string) string {
	suf := value[len(value)-2:] //px or dp and sp
	val := value[:len(value)-2]
	key := ""
	if suf == "sp" {
		key += fmt.Sprintf("text_size_%s_%s", suf, strings.Replace(val, ".", "_", -1))
	} else {
		key += fmt.Sprintf("ui_size_%s_%s", suf, strings.Replace(val, ".", "_", -1))
	}
	return key
}

func isHardCodeDimen(str string) bool {
	if str == "match_parent" || str == "wrap_content" || str == "fill_parent" || strings.HasPrefix(str, "@dimen/") ||
		strings.HasPrefix(str, "?") {
		return false
	}
	return true
}
