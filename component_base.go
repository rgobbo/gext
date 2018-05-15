package gext

type componentBase struct {
	id        string
	tagName   string
	children  []component
	src       string
	parent    component
	lastChild bool
	tagClose  bool
}

// Base returns the element base.
func (e *componentBase) Base() *componentBase {
	return e
}

// AppendChild appends the child element to the element.
func (e *componentBase) AppendChild(child component) {
	e.children = append(e.children, child)
}

//GetID object
func (e *componentBase) GetID() string {
	return e.id
}

//GetChildren : get children objects
func (e *componentBase) GetChildren() []component {
	return e.children
}

//GetChil : get children objects
func (e *componentBase) GetChild(tagname string) component {
	for _, o := range e.children {
		if o.GetTagName() == tagname {
			return o
		}
	}
	return nil
}

//GetChildById : get children objects
func (e *componentBase) GetChildById(id string) component {
	for _, o := range e.children {
		if o.GetID() == id {
			return o
		}
	}
	return nil
}

//GetTagName : get tagname
func (e *componentBase) GetTagName() string {
	return e.tagName
}
