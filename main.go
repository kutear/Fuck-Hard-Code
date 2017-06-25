package main

import (
	"flag"
	"github.com/kutear/fuck-hard-code/traversal"
	"log"
	"strings"
)

// 程序说明
// 目的:  解决Android项目中遗留的硬编码问题
// 参数:  layout:Android 项目中的layout目录
//		 config: json文件,内部表示具体哪些字段需要被替换
//		 existPixels:通常为/values/dimens.xml
//       existStrings:通常为/values/strings.xml
//		 scaleRatio:dp转化px的比例值 默认为3
//       out:输出修改后layout目录,注意不要与输入layout的一样
func main() {
	layoutPath := flag.String("layout", "", "the path of $project/appmodule/src/main/res/layout")
	configJson := flag.String("config", "", "configure file")
	existPixels := flag.String("existPixels", "", "the path of $project/appmodule/src/main/res/values/dimens.xml")
	existStrings := flag.String("existStrings", "", "the path of $project/appmodule/src/main/res/values/strings.xml")
	scaleRatio := flag.Float64("scaleRatio", 3, "px to dp,default 3px == 1dp")
	out := flag.String("out", "", "layout dir out path")
	flag.Parse()
	checkArgs(*layoutPath, "layout")
	checkArgs(*out, "out")
	traversal.PreTraversal(*configJson, *existPixels, *existStrings, *scaleRatio, *out)
	traversal.TraversalFile(*layoutPath)
}

func checkArgs(str, witch string) {
	if len(strings.TrimSpace(str)) == 0 {
		log.Fatalf("Args %s must be not Empty,Use -h to see all usage", witch)
	}
}
