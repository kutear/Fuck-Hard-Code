package parser

/**
 * 用来修改字符串硬编码
 */
type StringComponent struct {
	Text string `xml:"text,attr"`
	Hint string `xml:"hint,attr"`
	Sub  []StringComponent
}
