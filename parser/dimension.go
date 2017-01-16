package parser


/**
 * 用于修改尺寸硬编码
 */
type DimensionParser struct {
	Width  string `xml:"android:layout_width,attr"`
	Height string `xml:"android:layout_height,attr"`
}