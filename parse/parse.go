package parse

import "github.com/kutear/fuck-hard-code/parse/xml"

type Parse interface {
	//解析已经存在的属性值  string.xml dimens.xml
	ParseDefault(path string)
	//初始化一些信息 如dimens相关的radio
	InitExtra(extra interface{})
	//装载待处理的属性
	DealAttrs(attr interface{})
	//查看属性是否可以被处理
	Match(spaceName string, attrsName string) bool
	//修改Attrs
	Modify(attr *xml.Attr)
	//Save
	Save(path string)
}
