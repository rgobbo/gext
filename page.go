package gext

import (
	"bytes"
	"io"
	"net/http"
	"sync"
)

var varsMutex sync.Mutex

type eventPage struct {
	objectID  string
	eventType string
	funcName  string
}

type page struct {
	componentBase
	isPost       bool
	textValue    string
	pageVars     map[string]interface{}
	pageVarsInit map[string]interface{}
	Request      *http.Request
	MasterPage   string
	Controller   string
	Methods      []string
	Roles        []string
	Security     string
}

func (p *page) Execute(wr io.Writer, data map[string]interface{}, req *http.Request) error {
	if data != nil {
		p.pageVars = data
	}

	wr.Write([]byte(p.Render()))
	return nil
}

func (c *page) Render() string {
	var buffer bytes.Buffer

	if c.MasterPage != "" {
		master, _ := getPage(c.MasterPage)
		for _, c1 := range master.GetChildren() {
			if c1.GetTagName() == "block" {
				bl := c.GetChildById(c1.GetID())
				if bl != nil {
					buffer.WriteString(bl.Render())
				} else {
					buffer.WriteString(c1.Render())
				}
			} else {
				buffer.WriteString(c1.Render())
			}
		}

	} else {
		for _, c1 := range c.GetChildren() {
			buffer.WriteString(c1.Render())
		}
	}

	//   log.Println("html=",buffer.String())
	return buffer.String()
}
