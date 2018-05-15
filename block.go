package gext

import "bytes"

type block struct {
	componentBase
}

func (r *block) Render() string {
	var buffer bytes.Buffer
	for _, c1 := range r.GetChildren() {
		buffer.WriteString(c1.Render())
	}
	return buffer.String()
}

func newBlock(args map[string]interface{}) *block {
	o := &block{}
	o.id = args["id"].(string)
	o.tagName = "block"
	//	log.Println("Block created=", o.id)
	return o
}
