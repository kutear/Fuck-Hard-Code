package parser

/**
 * 用于修改尺寸硬编码
 */
type DimensionComponent struct {
	Width  string `xml:"layout_width,attr"`
	Height string `xml:"layout_height,attr"`
}
