package ssg

import (
	"fmt"
	"html"
	"strings"
)

// Viewable is an interface for renderable components.
type Viewable interface {
	Render() string
}

// HTMLString represents a raw HTML string.
type HTMLString string

// HTMLProps represents a map of HTML attributes.
type HTMLProps map[string]string

// Render returns the HTML string.
func (h HTMLString) Render() string {
	return string(h)
}

// HTMLTemplate represents a template with a format string and data.
type HTMLTemplate struct {
	Template string
	Data     []Viewable
}

// Render renders the template with the provided data.
func (h HTMLTemplate) Render() string {
	var renderedComponents []string
	for _, component := range h.Data {
		renderedComponents = append(renderedComponents, component.Render())
	}
	return fmt.Sprintf(h.Template, renderedComponents)
}

// HTMLComponent represents an HTML element with props and children.
type HTMLComponent struct {
	Name     string
	Props    HTMLProps
	Children []Viewable
}

// Render renders the HTML component with its props and children.
func (h HTMLComponent) Render() string {
	propsStr := h.Props.String()
	childrenStr := ""
	for _, child := range h.Children {
		childrenStr += child.Render()
	}
	return fmt.Sprintf("<%s %s>%s</%s>", h.Name, propsStr, childrenStr, h.Name)
}

// String converts HTMLProps to a string representation.
func (p HTMLProps) String() string {
	var propsStr []string
	for key, value := range p {
		propsStr = append(propsStr, fmt.Sprintf("%s=\"%s\"", key, html.EscapeString(value)))
	}
	return strings.Join(propsStr, " ")
}
