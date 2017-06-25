package parse

import (
	"encoding/xml"
)

//Strings.xml中的节点信息
type XmlString struct {
	XMLName xml.Name `xml:"resources"`
	Tags    []Node   `xml:"string"`
}
