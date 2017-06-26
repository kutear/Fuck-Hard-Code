package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//一个ViewGroup/View
type Element struct {
	Name   Name
	Attrs  []*Attr
	Child  []*Element
	Parent *Element
	Depth  int //在XmlRoot中的深度
}

//一个xml文件
type XmlRoot struct {
	Header  string
	Element *Element
	ns      map[string]string
}

//属性
type Attr struct {
	Value string
	Name  Name
}

//key值 例如 android:layout_width
type Name struct {
	xml.Name
	ShortSpace string
}

//解析文件xml
func ParseFile(path string) *XmlRoot {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("file %s parse faild", path)
	}
	return ParseBytes(bytes)
}

//解析[]byte
func ParseBytes(bytesData []byte) *XmlRoot {
	decoder := xml.NewDecoder(bytes.NewReader(bytesData))
	current := &XmlRoot{
		ns:     make(map[string]string, 0),
		Header: "",
	}
	for t, er := decoder.Token(); er == nil; t, er = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			parseStart(token, current)
		case xml.EndElement:
			parseEnd(token, current)
		case xml.CharData:
		case xml.Comment:
		case xml.Directive:
		case xml.ProcInst:
			current.Header = fmt.Sprint(`<?`, token.Copy().Target, ` `, string(token.Copy().Inst), `?>`)
		}
	}
	return current
}

// XmlRoot 格式化输出
func (this *XmlRoot) String() string {
	return fmt.Sprintf("header:%s \nnode: \n%s", this.Header, this.Element)
}

// Element 格式化输出
func (this *Element) String() string {
	str := fmt.Sprintf("nodeName: %s \n", this.Name)
	for _, el := range this.Child {
		str += el.String()
	}
	return str + "\n"
}

func parseStart(token xml.StartElement, parent *XmlRoot) {
	el := new(Element)
	dealNs(token.Name, parent, token.Name.Local)
	el.Name = Name{
		Name:       token.Name,
		ShortSpace: nameSpace(parent, token.Name.Space),
	}
	el.Parent = parent.Element
	el.Attrs = parseAttrs(token.Attr, parent)
	if parent != nil && parent.Element != nil {
		el.Depth = parent.Element.Depth + 1
		parent.Element.Child = append(parent.Element.Child, el)
	}
	parent.Element = el
}

func dealNs(name xml.Name, parent *XmlRoot, value string) {
	if strings.HasPrefix(value, "http") {
		parent.ns[value] = name.Local
	}
}

func parseAttrs(Attrs []xml.Attr, parent *XmlRoot) []*Attr {
	ret := make([]*Attr, len(Attrs))
	for index, attr := range Attrs {
		dealNs(attr.Name, parent, attr.Value)
		ret[index] = &Attr{
			attr.Value,
			Name{
				Name:       attr.Name,
				ShortSpace: nameSpace(parent, attr.Name.Space),
			},
		}
	}
	return ret
}

func nameSpace(parent *XmlRoot, space string) string {
	if value, ok := parent.ns[space]; ok {
		return value
	}
	return space
}

func parseEnd(token xml.EndElement, parent *XmlRoot) {
	if parent != nil && parent.Element != nil && parent.Element.Parent != nil {
		parent.Element = parent.Element.Parent
	}
}

//转化为XML格式字符输出
func (this *XmlRoot) ToXML() string {
	return fmt.Sprintf("%s \n%s", this.Header, this.Element.ToXML())
}

func (this *Element) ToXML() string {
	ret := ""
	attrsCount := len(this.Attrs)
	//childsCount := len(this.Child)
	//名字
	ret += spaceBlank(this.Depth)
	if this.Name.ShortSpace != "" {
		ret += fmt.Sprintf(startNodeWithNsFormat(attrsCount != 0), this.Name.ShortSpace, this.Name.Local)
	} else {
		ret += fmt.Sprintf(startNodeFormat(attrsCount != 0), this.Name.Local)
	}
	//属性
	for index, att := range this.Attrs {
		ret += spaceBlank(this.Depth)
		if len(att.Name.ShortSpace) > 0 {
			ret += fmt.Sprintf(attrWithNsFormat(index == attrsCount-1), att.Name.ShortSpace, att.Name.Local, att.Value)
		} else {
			ret += fmt.Sprintf(attrFormat(index == attrsCount-1), att.Name.Local, att.Value)
		}
	}
	if len(this.Child) > 0 {
		ret += fmt.Sprint(">\n")
		//Child
		for _, child := range this.Child {
			ret += fmt.Sprint(child.ToXML())
		}
		ret += spaceBlank(this.Depth)
		if this.Name.ShortSpace != "" {
			ret += fmt.Sprintf("</%s:%s>\n", this.Name.ShortSpace, this.Name.Local)
		} else {
			ret += fmt.Sprintf("</%s>\n", this.Name.Local)
		}
	} else {
		ret += fmt.Sprint("/>\n")
	}

	return ret
}

func startNodeWithNsFormat(nextLine bool) string {
	if nextLine {
		return "<%s:%s\n"
	}
	return "<%s:%s"
}

func startNodeFormat(nextLine bool) string {
	if nextLine {
		return "<%s\n"
	}
	return "<%s"
}

func attrWithNsFormat(last bool) string {
	if last {
		return "\t\t%s:%s=\"%s\""
	}
	return "\t\t%s:%s=\"%s\"\n"
}
func attrFormat(last bool) string {
	if last {
		return "\t\t%s=\"%s\""
	}
	return "\t\t%s=\"%s\"\n"
}

func spaceBlank(depth int) string {
	ret := ""
	for i := 0; i < depth; i++ {
		ret += "\t"
	}
	return ret
}
