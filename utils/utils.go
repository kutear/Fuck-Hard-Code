package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var mAttrs map[string]bool

func init() {
	mAttrs = make(map[string]bool)
	//常用的像素值
	mAttrs["layout_width"] = true
	mAttrs["layout_height"] = true
	mAttrs["layout_margin"] = true
	mAttrs["layout_marginBottom"] = true
	mAttrs["layout_marginEnd"] = true
	mAttrs["layout_marginLeft"] = true
	mAttrs["layout_marginRight"] = true
	mAttrs["layout_marginStart"] = true
	mAttrs["layout_marginTop"] = true
	mAttrs["padding"] = true
	mAttrs["paddingBottom"] = true
	mAttrs["paddingEnd"] = true
	mAttrs["paddingLeft"] = true
	mAttrs["paddingRight"] = true
	mAttrs["paddingStart"] = true
	mAttrs["paddingTop"] = true

	//常用的字符值
	mAttrs["text"] = true
	mAttrs["hint"] = true
}

/**
 * 判断给定的输入是否是字符串硬编码
 * 在Android中非硬编码是指引用string.xml中的字符串
 * 即形如“@string/***”就是非硬编码
 */
func isHardCodeString(str string) bool {
	return !strings.HasPrefix(str, "@string/") || !strings.HasPrefix(str, "?")
}

/**
 * 判断给定的输入是否是像素值硬编码
 * 在Android中非硬编码是指引用dimens.xml中的像素值
 * 即形如“@dimen/***”就是非硬编码
 */
func isHardCodeDimen(str string) bool {
	if str == "martch_parent" || str == "warp_content" || str == "fill_parent" {
		return false
	}
	return !strings.HasPrefix(str, "@dimen/") || !strings.HasPrefix(str, "?")
}

/**
 * 对于指定的属性是否需要做检查
 */
func needToCheck(attr string) bool {
	return mAttrs[attr]
}

func DealElement(element *Element, file string) {
	attrs := element.Attrs
	for i := 0; i < len(attrs); i++ {
		if needToCheck(attrs[i].Name()) {
			attrs[i].Value = "change"
		}
	}
	fmt.Println(element.ToXML())

	ioutil.WriteFile(file, []byte(element.ToXML()), 0644)
}
