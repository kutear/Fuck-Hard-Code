package utils

import (
	"fmt"
	"io/ioutil"
	"os"
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
	mAttrs["textSize"] = true

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
	if strings.HasPrefix(str, "@string/") || strings.HasPrefix(str, "?") {
		return false
	}
	return true
}

/**
 * 判断给定的输入是否是像素值硬编码
 * 在Android中非硬编码是指引用dimens.xml中的像素值
 * 即形如“@dimen/***”就是非硬编码
 */
func isHardCodeDimen(str string) bool {
	if str == "match_parent" || str == "wrap_content" || str == "fill_parent" || strings.HasPrefix(str, "@dimen/") ||
		strings.HasPrefix(str, "?") {
		return false
	}
	return true
}

/**
 * 对于指定的属性是否需要做检查
 */
func needToCheck(attr string) bool {
	return mAttrs[attr]
}

/**
 * 处理当页面
 */
func modifyElement(element *Element) {
	attrs := element.Attrs
	for i := 0; i < len(attrs); i++ {
		if needToCheck(attrs[i].Name()) {
			if attrs[i].Name() == "text" || attrs[i].Name() == "hint" {
				modifyStringAttr(attrs[i])
			} else {
				modifyDimenAttr(attrs[i])
			}
		}
	}
	childs := element.AllNodes()
	for i := 0; i < len(childs); i++ {
		modifyElement(childs[i])
	}
}

/**
 * file 文件全路径
 * simple 文件名称
 * outpath 输出路径
 */
func DealFile(file string, simple string, outpath string) {
	xmlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		panic("读取文件失败")
	}
	element, err := LoadByXml(string(xmlFile))
	if err != nil {
		panic("File [" + file + "] is Xml File ?")
	}
	out := outpath + string(os.PathSeparator) + simple
	fmt.Println(file + " > " + out)

	//递归修改每个节点的值
	modifyElement(element)

	ioutil.WriteFile(out, []byte(element.ToXML()), 0644)

}

/**
 * path 不存在时创建目录
 */
func CreatePath(path string) {
	exist, _ := pathExists(path)
	if !exist {
		os.Mkdir(path, 0644)
	}

}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
