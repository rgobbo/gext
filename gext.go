package gext

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/rgobbo/fsmodify"
)

var (
	cachedMutex          sync.Mutex
	generalPathTemplates string
	DefaultConfig        = GextConfig{LeftDelimeter: "{{", RightDelimeter: "}}", PathTemplates: "./", WatchInterval: 0, LogModify: true}
)

// Engine the goforms template engine
type Gext struct {
	cachedTemplates map[string]*template.Template
	cachedPages     map[string]*page
	funcs           map[string]interface{}
	Config          GextConfig
}

type GextConfig struct {
	PathTemplates  string
	LeftDelimeter  string
	RightDelimeter string
	WatchInterval  int
	LogModify      bool
}

// New creates and returns a new template engine
func NewGext(config GextConfig) *Gext {

	if !strings.HasSuffix(config.PathTemplates, "/") {
		config.PathTemplates = config.PathTemplates + "/"
	}
	generalPathTemplates = config.PathTemplates
	gext := &Gext{cachedTemplates: make(map[string]*template.Template), cachedPages: make(map[string]*page), Config: config, funcs: getFuncions()}

	if config.WatchInterval > 0 {
		go fsmodify.NewWatcher(config.PathTemplates, "", config.WatchInterval, func(filename string) {
			//base := path.Base(filename)
			ext := path.Ext(filename)
			if ext == ".html" {
				//gext.TemplateRemove(base)
				gext.ClearTemplateCache()
			}
		})
		//go watchfy.NewWatcher([]string{config.PathTemplates}, config.LogModify, func(filename string) {
		//	//base := path.Base(filename)
		//	ext := path.Ext(filename)
		//	if ext == ".html" {
		//		//gext.TemplateRemove(base)
		//		gext.ClearTemplateCache()
		//	}
		//})
	}

	return gext
}

func (g *Gext) Render(w io.Writer, name string, data interface{}) error {

	t, err := g.GetTemplate(name)
	if err != nil {
		return fmt.Errorf("[ERROR] Template with name %s dir :%s , error:", name, g.Config.PathTemplates, err.Error())
	}

	return t.ExecuteTemplate(w, name, data)

}

func (g *Gext) AddFuncs(funcs map[string]interface{}) {
	for s, f := range funcs {
		g.funcs[s] = f
	}
}

//GetTemplate - get parsed template from cache or parse new and cache
func (g *Gext) GetTemplate(name string) (*template.Template, error) {

	if t, ok := g.cachedTemplates[name]; ok {
		return t, nil
	}

	pgNew := &page{}
	if pg, ok := g.cachedPages[name]; ok {
		pgNew = pg
	} else {
		p, err := getPage(name)
		if err != nil {
			return nil, err
		}
		pgNew = p
	}

	cachedMutex.Lock()
	defer cachedMutex.Unlock()

	tmpl, err := template.New(name).Delims(g.Config.LeftDelimeter, g.Config.RightDelimeter).Funcs(g.funcs).Parse(pgNew.Render())
	if err != nil {
		return nil, err
	}
	g.cachedPages[name] = pgNew
	g.cachedTemplates[name] = tmpl

	return tmpl, nil

}

func (g *Gext) GetParsedPage(name string) (*page, error) {
	if p, ok := g.cachedPages[name]; ok {
		return p, nil
	}

	pg, err := getPage(name)
	if err != nil {
		return nil, err
	}
	cachedMutex.Lock()
	g.cachedPages[name] = pg
	defer cachedMutex.Unlock()

	return pg, nil

}

//ClearTemplateCache - Clear cached templates
func (g *Gext) ClearTemplateCache() {
	cachedMutex.Lock()
	defer cachedMutex.Unlock()
	for k := range g.cachedTemplates {
		delete(g.cachedTemplates, k)
		delete(g.cachedPages, k)
	}
}

//TemplateRemove - template remove from cache
func (g *Gext) TemplateRemove(name string) {
	cachedMutex.Lock()
	delete(g.cachedTemplates, name)
	delete(g.cachedPages, name)
	defer cachedMutex.Unlock()
}

func getPage(name string) (*page, error) {

	file, err := os.Open(generalPathTemplates + name)
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
