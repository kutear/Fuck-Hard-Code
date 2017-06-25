package parse

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

//解析string.xml和dimens.xml各到一个map中
func ParseValue(path string, xmlType interface{}, values map[string](string)) {
	if path == "" || values == nil {
		return
	}
	f, r := os.Open(path)
	defer f.Close()
	if r != nil {
		log.Fatalf("can't open this file %s,is Exist ?", path)
	}
	b, r := ioutil.ReadAll(f)
	if r != nil {
		log.Fatalf("can't read this file %s,is Exist ?", path)
	}
	xml.Unmarshal(b, &xmlType)
	var node []Node
	switch data := xmlType.(type) {
	case *XmlDimens:
		node = data.Tags
	case *XmlString:
		node = data.Tags
	default:
		log.Fatalf("can't deal the type of %s", reflect.TypeOf(xmlType))
	}
	for _, tag := range node {
		values[tag.Value] = tag.Key
	}
}
