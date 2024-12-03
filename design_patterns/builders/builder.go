package designpatterns

import (
	"fmt"
	"log"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))

	return sb.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)
	return b
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: HtmlElement{
			name:     rootName,
			text:     "",
			elements: []HtmlElement{},
		},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func DemonstrationOfBuilderPattern() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("<p>")
	log.Println(sb.String())

	// now with a list example
	words := []string{"hello", "world"}
	sb.Reset()

	// <ul> <li>...</li> ... </ul>
	sb.WriteString("<ul>")
	for _, word := range words {
		sb.WriteString("<li>")
		sb.WriteString(word)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	log.Println(sb.String())

	// checking the actual builder pattern
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "word")

	log.Println("builder pattern:", b.String())

}
