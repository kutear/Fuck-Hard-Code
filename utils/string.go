package utils

import "strconv"

var stringElement *Element
var stringIndex int
var strmp map[string]string

func init() {
	stringElement = NewElement("resources", "")
	stringIndex = 0
	strmp = make(map[string]string)
}

func modifyStringAttr(attr *Attr) {
	if isHardCodeString(attr.Value) {
		temp := attr.Value
		//The Value is Exist
		if key, ok := strmp[temp]; ok {
			attr.Value = "@string/" + key
		} else {
			v := "strings_" + strconv.Itoa(stringIndex)
			attr.Value = "@string/" + v
			addStringItem(v, temp)
			stringIndex++
		}
	}
}

func addStringItem(name string, v string) {
	strmp[v] = name
	node := NewElement("string", v)
	node.AddAttr("name", name)
	stringElement.AddNode(node)
}

func GetString() *Element {
	stringElement.head = "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	return stringElement
}
