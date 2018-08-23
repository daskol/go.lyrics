package genius

import (
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type visitFunc func(node *html.Node) bool

type Extractor struct {
	reader io.Reader
	root   *html.Node
	node   *html.Node
	text   string
}

func NewExtractor(reader io.Reader) *Extractor {
	return &Extractor{reader: reader}
}

func (e *Extractor) Extract() (string, error) {
	if root, err := html.Parse(e.reader); err != nil {
		return "", err
	} else {
		e.root = root
		e.walk(e.root, e.findDivLyrics)
		e.walk(e.node, e.htmlToText)
		return e.text, nil
	}
}

func (e *Extractor) htmlToText(node *html.Node) bool {
	if node.Type == html.TextNode {
		e.text += node.Data
	}
	return true
}

func (e *Extractor) findDivLyrics(node *html.Node) bool {
	if node.DataAtom != atom.Div {
		return true
	}

	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "lyrics" {
			e.node = node
			return false
		}
	}

	return true
}

func (d *Extractor) walk(node *html.Node, fn visitFunc) {
	if node.Type == html.CommentNode ||
		node.Type == html.DoctypeNode ||
		node.Type == html.ErrorNode {
		return
	}

	if !fn(node) {
		return
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		d.walk(child, fn)
	}
}
