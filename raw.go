package gext

import "golang.org/x/net/html"

type raw struct {
	componentBase
	classes          []string
	containPlainText bool
	textValue        string
	attr             []html.Attribute
}

func (r *raw) Render() string {
	// var buffer bytes.Buffer
	// buffer.WriteString(r.textValue)
	// for _, c1 := range r.GetChildren() {
	// 	buffer.WriteString(c1.Render())
	// }
	// if r.Base().tagClose == true {
	// 	if r.tagName != "link" && r.tagName != "meta" {
	// 		buffer.WriteString("</")
	// 		buffer.WriteString(r.tagName)
	// 		buffer.WriteString(">")

	// 	}
	// }
	// return buffer.String()
	return r.textValue
}

func newRaw(lit string) *raw {
	o := &raw{}
	o.tagName = "html"
	o.textValue = lit // html.UnescapeString(lit)
	return o
}
