package htmls

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type (
	Node struct {
		// use struct member instead of "type Node html.Node", because need to hide methods of html.Node
		raw *html.Node
	}
	
	Nodes struct {
		items []*Node
	}

	HtmlDoc struct {
		node *html.Node
	}

	HtmlNodes struct {
		nodes []*html.Node
	}

	NodeEachCallback func(i int, n *Node)
)

func (n *Node) Attr(attrKey string) string {
	if n == nil {
		return ""
	}
	for _, v := range n.raw.Attr {
		if v.Key == attrKey {
			return v.Val
		}
	}
	return ""
}

func (n *Node) Text() string {
	if n == nil {
		return ""
	}
	return htmlquery.InnerText(n.raw)
}

func (n *Node) Query(expr string) *Nodes {
	if n == nil {
		return nil
	}
	ns := htmlquery.Find(n.raw, expr)
	res := &Nodes{}
	for _, v := range ns {
		if v == nil {
			continue
		}
		res.items = append(res.items, &Node{raw:v})
	}
	return res
}

func (n *Node) QueryByElementAndFullAttrVal(element, attrKey, fullAttrVal string) *Nodes {
	if n == nil {
		return nil
	}
	ns := htmlquery.Find(n.raw, "//"+element)
	res := &Nodes{}
	for _, v := range ns {
		if v == nil {
			continue
		}
		n := &Node{raw:v}
		if n.Attr(attrKey) == fullAttrVal {
			res.items = append(res.items, n)
		}
	}
	return res
}

func (n *Node) QueryByElementAndPartAttrVal(element, attrKey, partAttrVal string) *Nodes {
	if n == nil {
		return nil
	}
	ns := htmlquery.Find(n.raw, "//"+element)
	res := &Nodes{}
	for _, v := range ns {
		if v == nil {
			continue
		}
		n := &Node{raw:v}
		if strings.Contains(n.Attr(attrKey), partAttrVal) {
			res.items = append(res.items, n)
		}
	}
	return res
}

func (n *Node) Parent() *Node {
	return &Node{raw:n.raw.Parent}
}

func (ns *Nodes) Get(i int) *Node {
	if ns == nil || len(ns.items) == 0 || i > ns.Len() - 1 {
		return nil
	}
	return ns.items[i]
}

func (ns *Nodes) First() *Node {
	if ns == nil || len(ns.items) == 0 {
		return nil
	}
	return ns.items[0]
}

func (ns *Nodes) Last() *Node {
	if ns == nil || len(ns.items) == 0 {
		return nil
	}
	return ns.items[ns.Len() - 1]
}

func (ns *Nodes) Len() int {
	if ns == nil {
		return 0
	}
	return len(ns.items)
}

func (ns *Nodes) Each(cb NodeEachCallback) {
	for i := 0; i < ns.Len(); i++ {
		cb(i, ns.Get(i))
	}
}

func (ns *Nodes) Attrs(attrKey string) []string {
	var res []string
	for i := 0; i < ns.Len(); i++ {
		res = append(res, ns.Get(i).Attr(attrKey))
	}
	return res
}

func (ns *Nodes) Texts() []string {
	var res []string
	for i := 0; i < ns.Len(); i++ {
		res = append(res, ns.Get(i).Text())
	}
	return res
}

func NewFromString(s string) (*Node, error) {
	node, err := htmlquery.Parse(strings.NewReader(s))
	if err != nil {
		return nil, err
	}
	return &Node{raw:node}, nil
}