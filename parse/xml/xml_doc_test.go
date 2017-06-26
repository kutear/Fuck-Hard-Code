package xml

import (
	"strings"
	"testing"
)

var xmlStr = `
		<?xml version="1.0" encoding="utf-8"?>
		<X name="test">
			<Y desc="this_is_y"/>
			<Y desc="this_is_y"/>
			<Y desc="this_is_y"/>
			<Y desc="this_is_y"/>
		</X>
	`

func TestParseBytes(t *testing.T) {
	root := ParseBytes([]byte(xmlStr))
	if root.Header != "<?xml version=\"1.0\" encoding=\"utf-8\"?>" {
		t.Error("Header Error")
	}

	if len(root.Element.Child) != 4 {
		t.Error("Child Num Error")
	}
}

func TestXmlRoot_ToXML(t *testing.T) {
	root := ParseBytes([]byte(xmlStr))
	noSpace := strings.Replace(xmlStr, " ", "", -1)
	noSpace = strings.Replace(noSpace, "\t", "", -1)
	noSpace = strings.Replace(noSpace, "\n", "", -1)
	after := strings.Replace(root.ToXML(), " ", "", -1)
	after = strings.Replace(after, "\t", "", -1)
	after = strings.Replace(after, "\n", "", -1)
	if noSpace != after {
		t.Error("to xml error")
	}
}
