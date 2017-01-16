package utils

import "strings"


/**
 * 判断给定的输入是否是字符串硬编码
 * 在Android中非硬编码是指引用string.xml中的字符串
 * 即形如“@string/***”就是非硬编码
 */
func isHardCodeString(str string) bool {
	return !strings.HasPrefix(str, "@string/")
}


/**
 * 判断给定的输入是否是像素值硬编码
 * 在Android中非硬编码是指引用dimens.xml中的像素值
 * 即形如“@dimen/***”就是非硬编码
 */
func isHardCodeDimen(str string) bool {
	return !strings.HasPrefix(str, "@dimen/")
}
