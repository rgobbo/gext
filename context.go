package gext

import (
	"net/http"
)

// context Request context
type context struct {
	PageObj *page
	Data    map[string]interface{}
	Request *http.Request
}

// SetValue get label component by id from page
func (ctx *context) SetValue(obj string, prop string, value interface{}) {
	ctx.Data[obj+prop] = value

}
