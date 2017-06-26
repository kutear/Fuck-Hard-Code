package parse

import (
	"encoding/xml"
)

//dimens.xml中的节点信息
type XmlDimens struct {
	XMLName xml.Name `xml:"resources"`
	Tags    []Node   `xml:"dimen"`
}
