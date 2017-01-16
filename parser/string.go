package parser

/**
 * 用来修改字符串硬编码
 */
type StringParser struct {
	Text string `xml:"android:text,attr"`
	Hint string `xml:"android:hint,attr"`
}

