package utils

import (
	"github.com/kutear/fuck-hard-code/parse"
	"io/ioutil"
	"os"
	"sort"
)

//创建路径
func CreatePath(path string) string {
	exist, _ := pathExists(path)
	if !exist {
		os.Mkdir(path, 0755)
	}
	return path
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

//检查数组中是否包含某个string
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//文件写入
func Write(data []byte, file string) {
	ioutil.WriteFile(file, data, 0755)
}

//对map进行排序输出
func SortMap(sortMap map[string]string) []parse.Node {
	keys := make([]string, len(sortMap))
	i := 0
	for k := range sortMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	ret := make([]parse.Node, len(sortMap))
	i = 0
	for _, v := range keys {
		ret[i] = parse.Node{
			Value: v,
			Key:   sortMap[v],
		}
		i++
	}
	return ret
}
