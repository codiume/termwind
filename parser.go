package termwind

import (
	"strings"

	"golang.org/x/net/html"
)

var supportedTags = map[Tag]bool{
	TagAnchor:    true,
	TagBreakLine: true,
	TagDd:        true,
	TagDiv:       true,
	TagDl:        true,
	TagDt:        true,
	TagHr:        true,
	TagLi:        true,
	TagOl:        true,
	TagParagraph: true,
	TagRaw:       true,
	TagSpan:      true,
	TagUl:        true,
}

func Parse(input string) (*Element, error) {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "<") {
		input = "<div>" + input + "</div>"
	}

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return nil, err
	}

	root := findContent(doc)
	if root == nil {
		return &Element{Tag: TagDiv}, nil
	}

	return root, nil
}

func findContent(node *html.Node) *Element {
	if node.Type == html.ElementNode {
		tag := Tag(node.Data)
		if supportedTags[tag] {
			return walkNode(node, nil)
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if el := findContent(child); el != nil {
			return el
		}
	}

	return nil
}

func walkNode(node *html.Node, parent *Element) *Element {
	if node.Type == html.TextNode {
		if parent == nil {
			return nil
		}
		text := node.Data
		if !parent.Tag.IsInline() {
			text = strings.TrimSpace(node.Data)
		}
		if text == "" {
			return nil
		}
		textEl := &Element{Tag: TagRaw, Content: text, Parent: parent}
		parent.Children = append(parent.Children, textEl)
		return nil
	}

	if node.Type != html.ElementNode {
		walkChildren(node, parent)
		return nil
	}

	tag := Tag(node.Data)
	if !supportedTags[tag] {
		walkChildren(node, parent)
		return nil
	}

	el := &Element{
		Tag:    tag,
		Parent: parent,
	}

	for _, attr := range node.Attr {
		if attr.Key == "class" {
			el.Classes = strings.Fields(attr.Val)
		} else {
			if el.Attrs == nil {
				el.Attrs = make(map[string]string)
			}
			el.Attrs[attr.Key] = attr.Val
		}
	}

	if len(el.Classes) > 0 {
		el.Style = ParseClasses(el.Classes)
	}

	walkChildren(node, el)

	return el
}

// walkChildren recurses through all siblings of a node's children,
// appending any resulting elements to the parent's Children slice.
func walkChildren(node *html.Node, parent *Element) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		child := walkNode(c, parent)
		if child != nil && parent != nil {
			parent.Children = append(parent.Children, child)
		}
	}
}
