package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var dimenElement *Element
var kv map[string]string

func init() {
	dimenElement = NewElement("resources", "")
	kv = make(map[string]string)
}

func modifyDimenAttr(attr *Attr) {
	value := attr.Value
	if isHardCodeDimen(value) {
		//because dip == dp
		if strings.HasSuffix(value, "dip") {
			value = strings.Replace(value, "dip", "dp", -1)
		}

		suf := value[len(value)-2:] //px or dp and sp
		temp := getDimenFromValue(value)
		res := fmt.Sprintf("%.1f", temp)
		if suf == "px" {
			suf = "dp"
		}
		key := suf + "_" + fmt.Sprintf("%06.1f", temp)
		key = strings.Replace(key, ".", "_", -1)
		key = strings.Replace(key, "-", "n_", -1)
		_, ok := kv[key]
		if !ok {
			addDimenItem(key, res+suf)
		}
		attr.Value = "@dimen/" + key
	}
}

func getDimenFromValue(v string) float64 {
	value := 0.0
	temp := v[0 : len(v)-2]
	f, e := strconv.ParseFloat(temp, 64)
	if e != nil {
		fmt.Println(e.Error())
	}
	if strings.HasSuffix(v, "px") {
		value = px2dp(f)
	} else {
		value = f
	}
	return value
}

func px2dp(px float64) float64 {
	return px / 3
}

func addDimenItem(key, v string) {
	kv[key] = v
}

func GetDimen() *Element {
	dpmap := make(map[string]string)
	for k, v := range kv {
		if strings.HasPrefix(k, "dp") {
			dpmap[k] = v
		}
	}
	sortInsert(dpmap)
	spmap := make(map[string]string)
	for k, v := range kv {
		if strings.HasPrefix(k, "sp") {
			spmap[k] = v
		}
	}
	sortInsert(spmap)

	dimenElement.head = "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	return dimenElement
}

func sortInsert(mp map[string]string) {
	arr := make([]string, 0)

	for k := range mp {
		arr = append(arr, k)
	}

	sort.Strings(arr)

	for i := 0; i < len(mp); i++ {
		v := mp[arr[i]]
		buildItem(arr[i], v)
	}
}

func buildItem(key, v string) {
	node := NewElement("dimen", v)
	node.AddAttr("name", key)
	dimenElement.AddNode(node)
}
