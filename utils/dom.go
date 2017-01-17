/**
 *donnie4w@gmail.com
 */
package utils

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"runtime/debug"
	"strings"
	"sync"
)

const _VAR = "1.0.2"

type E interface {
	ToString() string
}

type Attr struct {
	space string
	name  string
	Value string
}

func (a *Attr) Name() string {
	return a.name
}

type Element struct {
	head       string
	space      string
	name       string
	Value      string
	Attrs      []*Attr
	childs     []E
	parent     E
	elementmap map[string][]E
	attrmap    map[string]string
	lc         *sync.RWMutex
	r          E
	root       E
	isSync     bool
}

func LoadByStream(r io.Reader) (current *Element, err error) {
	defer func() {
		if er := recover(); er != nil {
			fmt.Println(er)
			fmt.Println(string(debug.Stack()))
			err = errors.New("xml load error!")
		}
	}()
	decoder := xml.NewDecoder(r)
	isRoot := true
	head := ""
	for t, er := decoder.Token(); er == nil; t, er = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			el := new(Element)
			el.space = space(token.Name.Space)
			el.name = token.Name.Local
			el.Attrs = make([]*Attr, 0)
			el.childs = make([]E, 0)
			el.elementmap = make(map[string][]E, 0)
			el.attrmap = make(map[string]string, 0)
			el.lc = new(sync.RWMutex)
			el.r = el
			el.isSync = false
			for _, a := range token.Attr {
				ar := new(Attr)
				ar.space = space(a.Name.Space)
				ar.name = a.Name.Local
				ar.Value = a.Value
				el.Attrs = append(el.Attrs, ar)
				el.attrmap[ar.name] = ar.Value
			}
			if isRoot {
				isRoot = false
				el.root = el
			} else {
				current.childs = append(current.childs, el)
				current.elementmap[el.name] = append(current.elementmap[el.name], el)
				el.parent = current
				el.root = current.root
			}
			current = el
		case xml.EndElement:
			if current.parent != nil {
				current = current.parent.(*Element)
			}
		case xml.CharData:
			if token != nil && current != nil {
				current.Value = string([]byte(token.Copy()))
			}
		case xml.Comment:
		//			fmt.Println("xml===>1", string(token.Copy()))
		case xml.Directive:
		//			fmt.Println("xml===>2", string(token.Copy()))
		case xml.ProcInst:
			head = fmt.Sprint(`<?`, token.Copy().Target, ` `, string(token.Copy().Inst), `?>`)
		default:
			panic("parse xml fail!")
		}
	}
	current.Root().head = head
	return current, nil
}

func LoadByXml(xmlstr string) (current *Element, err error) {
	defer func() {
		if er := recover(); er != nil {
			fmt.Println(er)
			err = errors.New("xml load error!")
		}
	}()
	s := strings.NewReader(xmlstr)
	return LoadByStream(s)
}

func (t *Element) ToString() string {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	return t._string()
}

func (t *Element) Name() string {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	}
	return t.name
}

func (t *Element) Head() string {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	}
	return t.head
}

func NewElement(elementName, elementValue string) (el *Element) {
	el = &Element{name: elementName, Value: elementValue, Attrs: make([]*Attr, 0), childs: make([]E, 0), elementmap: make(map[string][]E, 0), attrmap: make(map[string]string, 0), lc: new(sync.RWMutex), isSync: false}
	el.root = el
	el.r = el
	return
}

func (t *Element) _string() string {
	elementname := t.name
	if t.space != "" {
		elementname = fmt.Sprint(t.space, ":", elementname)
	}
	s := fmt.Sprint("<", elementname)
	sattr := ""
	if len(t.Attrs) > 0 {
		for _, att := range t.Attrs {
			attrname := att.name
			if att.space != "" {
				attrname = fmt.Sprint(att.space, ":", attrname)
			}
			sattr = fmt.Sprint(sattr, " ", attrname, "=", "\"", att.Value, "\"")
		}
	}
	s = fmt.Sprint(s, sattr, ">")
	if len(t.childs) > 0 {
		for _, v := range t.childs {
			el := v.(*Element)
			s = fmt.Sprint(s, el._string())
		}
		return fmt.Sprint(s, t.Value, "</", elementname, ">", "\n")
	} else {
		return toStr(t)
	}
}

func toStr(t *Element) string {
	sattr := ""
	if len(t.Attrs) > 0 {
		for _, att := range t.Attrs {
			attrname := att.name
			if att.space != "" {
				attrname = fmt.Sprint(att.space, ":", attrname)
			}
			sattr = fmt.Sprint(sattr, " ", attrname, "=", "\"", att.Value, "\"")
		}
	}
	return fmt.Sprint("<", t.name, sattr, ">", t.Value, "</", t.name, ">")
}

//return child element "name"
func (t *Element) Node(name string) *Element {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	es, ok := t.elementmap[name]
	if ok {
		el := es[0]
		return el.(*Element)
	} else {
		return nil
	}
}

func (t *Element) GetNodeByPath(path string) *Element {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	paths := strings.Split(path, "/")
	if paths != nil && len(paths) > 0 {
		e := t
		for i, p := range paths {
			if i == 0 {
				if e.Name() == p {
					continue
				} else {
					return nil
				}
			}
			e = e.Node(p)
			if e == nil {
				return nil
			}
		}
		return e
	}
	return nil
}

func (t *Element) GetNodesByPath(path string) []*Element {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	paths := strings.Split(path, "/")
	if paths != nil {
		length := len(paths)
		if length > 0 {
			if length == 1 {
				return t.Nodes(paths[0])
			}
			d_name := paths[length-1]
			d_name_len := len(d_name)
			sup_nodepath := path[:length-d_name_len]
			sup_node := t.GetNodeByPath(sup_nodepath)
			return sup_node.Nodes(d_name)
		}
	}
	return nil
}

// return child element length
func (t *Element) NodesLength() int64 {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	if t.childs != nil {
		return int64(len(t.childs))
	} else {
		return 0
	}
}

// whole xml length
func (t *Element) DocLength() int64 {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	var retc int64
	for _, v := range t._root().childs {
		el := v.(*Element)
		retc = retc + el._elementLen()
	}
	return retc + 1
}

func (t *Element) _elementLen() int64 {
	if len(t.childs) > 0 {
		var retc int64
		for _, v := range t.childs {
			el := v.(*Element)
			retc = retc + el._elementLen()
		}
		return retc + 1
	} else {
		return 1
	}
}

// return all the child element "name"
func (t *Element) Nodes(name string) []*Element {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	es, ok := t.elementmap[name]
	if ok {
		ret := make([]*Element, len(es))
		for i, v := range es {
			ret[i] = v.(*Element)
		}
		return ret
	} else {
		return nil
	}
}

func (t *Element) AttrValue(name string) (string, bool) {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	v, ok := t.attrmap[name]
	if ok {
		return v, true
	} else {
		return "", false
	}
}

func (t *Element) AddAttr(name, value string) {
	if t._root().isSync {
		t._root().lc.Lock()
		defer t._root().lc.Unlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.Lock()
		defer rt.lc.Unlock()
	}
	t.attrmap[name] = value
	isExist := false
	for _, v := range t.Attrs {
		if v.name == name {
			v.Value = value
			isExist = true
		}
	}
	if !isExist {
		t.Attrs = append(t.Attrs, &Attr{"", name, value})
	}
}

//remove the attribute "name" for current element
func (t *Element) RemoveAttr(name string) bool {
	if t._root().isSync {
		t._root().lc.Lock()
		defer t._root().lc.Unlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.Lock()
		defer rt.lc.Unlock()
	}
	_, ok := t.attrmap[name]
	if ok {
		delete(t.attrmap, name)
		newAs := make([]*Attr, 0)
		for _, v := range t.Attrs {
			if v.name != name {
				newAs = append(newAs, v)
			}
		}
		t.Attrs = newAs
		return true
	} else {
		return false
	}
}

//return all the child elements
func (t *Element) AllNodes() []*Element {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	es := t.childs
	if len(es) > 0 {
		ret := make([]*Element, len(es))
		for i, v := range es {
			ret[i] = v.(*Element)
		}
		return ret
	} else {
		return nil
	}
}

//remove the child element "name" for current element
func (t *Element) RemoveNode(name string) bool {
	if t._root().isSync {
		t._root().lc.Lock()
		defer t._root().lc.Unlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.Lock()
		defer rt.lc.Unlock()
	}
	_, ok := t.elementmap[name]
	if ok {
		delete(t.elementmap, name)
		newCs := make([]E, 0)
		for _, v := range t.childs {
			if v.(*Element).name != name {
				newCs = append(newCs, v)
			}
		}
		t.childs = newCs
		return true
	} else {
		return false
	}
}

// return the root element
func (t *Element) Root() *Element {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	return t._root()
}

func (t *Element) _root() *Element {
	return t.root.(*Element)
}

func (t *Element) AddNode(el *Element) error {
	if t._root().isSync {
		t._root().lc.Lock()
		defer t._root().lc.Unlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.Lock()
		defer rt.lc.Unlock()
	}
	return t._addNode(el)
}

func (t *Element) _addNode(el *Element) error {
	if el.name == "" {
		return errors.New("error!|name is empty!")
	}
	t.childs = append(t.childs, el)
	el.parent = t
	el.r = el
	el.changeRoot(t._root())
	t.elementmap[el.name] = append(t.elementmap[el.name], el)
	return nil
}

func (t *Element) changeRoot(el *Element) {
	if len(t.childs) > 0 {
		for _, v := range t.childs {
			v.(*Element).changeRoot(el)
		}
	}
	t.root = el
}

//add an element used string for current element
func (t *Element) AddNodeByString(xmlstr string) error {
	if t._root().isSync {
		t._root().lc.Lock()
		defer t._root().lc.Unlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.Lock()
		defer rt.lc.Unlock()
	}
	el, err := LoadByXml(xmlstr)
	if err != nil {
		return err
	}
	t._addNode(el)
	return nil
}

//current element's parent
func (t *Element) Parent() *Element {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	if t.parent != nil {
		return t.parent.(*Element)
	} else {
		return nil
	}
}

//whole xml
func (t *Element) ToXML() string {
	if t._root().isSync {
		t._root().lc.RLock()
		defer t._root().lc.RUnlock()
	} else {
		rt := t.r.(*Element)
		rt.lc.RLock()
		defer rt.lc.RUnlock()
	}
	return t._root()._string()
}

func (t *Element) SyncToXml() string {
	t._root().isSync = true
	t._root().lc.Lock()
	defer func() {
		t._root().lc.Unlock()
		t._root().isSync = false
	}()
	return t._root()._string()
}

func space(spacename string) string {
	i := strings.LastIndex(spacename, "/")
	if i > 0 {
		spacename = spacename[i+1:]
	}
	return spacename
}
