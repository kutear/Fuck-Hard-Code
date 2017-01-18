package utils

import "strconv"

var stringElement *Element
var stringIndex int

func init() {
	stringElement = NewElement("resources", "")
	stringIndex = 0
}

func modifyStringAttr(attr *Attr) {
	if isHardCodeString(attr.Value) {
		temp := attr.Value
		v := "strings_" + strconv.Itoa(stringIndex)
		attr.Value = "@string/" + v
		addStringItem(v, temp)
		stringIndex++
	}
}

func addStringItem(name string, v string) {
	node := NewElement("string", v)
	node.AddAttr("name", name)
	dimenIndex++
	stringElement.AddNode(node)
}

func GetString() *Element {
	return stringElement
}
