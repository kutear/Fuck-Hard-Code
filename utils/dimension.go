package utils

import (
	"fmt"
	"strconv"
	"strings"
)

var dimenElement *Element
var dimenIndex int

func init() {
	dimenElement = NewElement("resources", "")
	dimenIndex = 0
}

func modifyDimenAttr(attr *Attr) {
	value := attr.Value
	if isHardCodeDimen(value) {
		suf := value[len(value)-2:]
		temp := getDimenFromValue(value)
		key := "dp_" + strconv.Itoa(dimenIndex)
		attr.Value = "@dimen/" + key
		res := fmt.Sprintf("%.1f", temp)
		if suf == "px" {
			suf = "dp"
		}
		addDimenItem(key, res+suf)
	}
}

func getDimenFromValue(v string) float64 {
	value := 0.0
	temp := v[0 : len(v)-2]
	f, e := strconv.ParseFloat(temp, 64)
	if e != nil {
		fmt.Println(e.Error())
	}
	if strings.HasSuffix(v, "px") {
		value = px2dp(f)
	} else {
		value = f
	}
	return value
}

func px2dp(px float64) float64 {
	return px / 3
}

func addDimenItem(key, v string) {
	node := NewElement("dimen", v)
	node.AddAttr("name", key)
	dimenIndex++
	dimenElement.AddNode(node)
}

func GetDimen() *Element {
	return dimenElement
}
