package gext

type component interface {
	Render() string
	AppendChild(component)
	GetID() string
	GetChildren() []component
	GetTagName() string
	Base() *componentBase
}

// func newComponent(token html.Token, urlpath string) (component, map[string]interface{}, error) {
// 	var e component
// 	vars := make(map[string]interface{})

// 	switch token.Data {
// 	case "go:button":
// 		e = newButton(token)
// 	case "go:label":
// 		e, vars = newLabel(token)
// 	case "go:form":
// 		e = newForm(token, urlpath)
// 	case "go:dropdown":
// 		e = newDropDown(token)
// 	case "go:listitem":
// 		e = newListItem(token)
// 	case "go:extends":
// 		e, vars = newMaster(token)
// 	case "go:block":
// 		e = newBlock(token)
// 	case "go:repeater":
// 		e = newRepeater(token)
// 	case "go:itemtemplate":
// 		e = newItemTemplate(token)
// 	case "go:headertemplate":
// 		e = newHeaderTemplate(token)
// 	case "go:footertemplate":
// 		e = newFooterTemplate(token)
// 	case "go:textbox":
// 		e, vars = newTextBox(token)
// 	default:
// 		e = newRaw(token)
// 	}

// 	return e, vars, nil
// }
