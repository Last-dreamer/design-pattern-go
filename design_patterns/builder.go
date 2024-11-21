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
	sb.WriteString(fmt.Sprintf("%d<%s>\n",
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
	return sb.String()
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

}
