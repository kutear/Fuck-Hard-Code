package parse

//XML <xxx name="yyy"> zzz </xxx>
type Node struct {
	Key   string `xml:"name,attr"`
	Value string `xml:",chardata"`
}
