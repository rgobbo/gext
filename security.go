package gext

type security struct {
	componentBase
	description string
}

func (s *security) Render() string {
	return ""
}

func newSecurity(desc string) *security {
	sec := &security{}
	//filename := args["filename"].(string)
	//name := strings.Replace(filename, ".html", "", 1)

	sec.description = desc
	return sec
}
