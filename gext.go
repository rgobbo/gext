package gext

import (
	"html/template"
	"io"
	"os"
	"path"
	"sync"
	"github.com/rgobbo/watchfy"

	"fmt"
	"strings"
)

var (
	cachedMutex sync.Mutex
    generalPathTemplates string
    DefaultConfig = GextConfig{LeftDelimeter:"{{",RightDelimeter:"}}",PathTemplates:"./",WatchModify:true, LogModify:true}
)


// Engine the goforms template engine
type Gext struct {
	cachedTemplates  map[string]*template.Template
	funcs map[string]interface{}
	Config GextConfig
}

type GextConfig struct {
	PathTemplates string
	LeftDelimeter  string
	RightDelimeter string
	WatchModify bool
	LogModify bool
}

// New creates and returns a new template engine
func NewGext(config GextConfig) *Gext {

	if !strings.HasSuffix(config.PathTemplates, "/") {
		config.PathTemplates = config.PathTemplates + "/"
	}
	generalPathTemplates = config.PathTemplates
	gext := &Gext{cachedTemplates:make(map[string]*template.Template), Config:config, funcs:getFuncions()}

	if config.WatchModify == true {
		go watchfy.NewWatcher([]string{config.PathTemplates}, config.LogModify, func(filename string) {
			//base := path.Base(filename)
			ext := path.Ext(filename)
			if ext == ".html" {
				//gext.TemplateRemove(base)
				gext.ClearTemplateCache()
			}
		})
	}

	return gext
}


func (g *Gext) Render(w io.Writer, name string, data interface{} ) error {

	t, err := g.GetTemplate(name)
	if err != nil {
		return fmt.Errorf("[NOT FOUND] Template with name %s doesn't exists in the dir :%s", name, g.Config.PathTemplates)
	}

	return t.ExecuteTemplate(w, name, data)

}

func (g *Gext)AddFuncs(funcs map[string]interface{}) {
	for s,f := range funcs {
		g.funcs[s] = f
	}
}

//GetTemplate - get parsed template from cache or parse new and cache
func (g *Gext) GetTemplate(name string) (*template.Template, error) {

	if t, ok := g.cachedTemplates[name]; ok {
		return t, nil
	}

	p, err := getPage(name)
	if err != nil {
		return nil, err
	}


	cachedMutex.Lock()
	tmpl := template.Must(template.New(name).Delims(g.Config.LeftDelimeter, g.Config.RightDelimeter).Funcs(g.funcs).Parse(p.Render()))
	if err != nil {
		return nil, err
	}
	g.cachedTemplates[name] = tmpl
	defer cachedMutex.Unlock()

	return tmpl, nil

}


//ClearTemplateCache - Clear cached templates
func (g *Gext) ClearTemplateCache() {
	cachedMutex.Lock()
	defer cachedMutex.Unlock()
	for k := range g.cachedTemplates {
		delete(g.cachedTemplates, k)
	}
}

//TemplateRemove - template remove from cache
func (g *Gext) TemplateRemove(name string) {
	cachedMutex.Lock()
	delete(g.cachedTemplates, name)
	defer cachedMutex.Unlock()
}



func getPage(name string) (*page, error) {

	file, err := os.Open( generalPathTemplates + name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	path := ""
	if name == "index.html" {
		path = "/"
	} else {
		path = name
	}

	parser := newParser(file)
	p, err := parser.Parse(path)
	if err != nil {
		return nil, err
	}

	p.id = name

	return p, nil

}