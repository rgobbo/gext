package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rgobbo/gext"
	"time"
	"log"
	"strings"
)

var engine *gext.Gext
func main() {
	//config := gext.GextConfig{
	//	                      PathTemplates:"./templates",
	//	                      LeftDelimeter: "{{",
	//	                      RightDelimeter: "}}",
	//	                      WatchModify: true,
	//	                      LogModify: false}
	// engine = gext.NewGext(config)
	// or use DefaultConfig
	engine = gext.NewGext(gext.DefaultConfig)

	// add new function to engine
	funcs := map[string] interface{}{
		"Count" : strings.Count,
	}
	engine.AddFuncs(funcs)


	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Println("Starting server : 127.0.0.1:8080")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := engine.Render(w,"index.html", nil)
	if err != nil {
		log.Println("Render error:", err)
	}
}