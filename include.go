package gext

import "log"

type include struct {
	componentBase
	includePage string
}

func (in *include) Render() string {
	p, err := getPage(in.includePage)
	if err != nil {
		log.Println("include render error:=", err)
	}
	return p.Render()
}

func newInclude(filename string) *include {
	in := &include{}
	//filename := args["filename"].(string)
	//name := strings.Replace(filename, ".html", "", 1)
	p, err := getPage(filename)
	if err != nil {
		log.Println("include error:=", err)
		return nil
	}
	p.tagName = "includefile"
	in.includePage = p.id
	return in
}
