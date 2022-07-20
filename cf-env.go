package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var bgColor = "white" // set a sane default color for background

type Environment struct {
	Color       string
	Environment []string
	Header      map[string][]string
}

func (e *Environment) SetColor(c string) {
	e.Color = c
}

func (e *Environment) SetEnvironment() {
	e.Environment = os.Environ()
}

func (e *Environment) SetHeader(r *http.Request) {
	e.Header = r.Header
}

func killHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(5)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
		env := Environment{}
		env.SetEnvironment()
		env.SetHeader(r)
		env.SetColor(bgColor)

		t := template.New("index.html")
		t, err := t.ParseFiles("tmpl/index.html")
		if err != nil {
			log.Fatal("ParseFiles: ", err)
		}
		if t == nil {
			log.Fatal("ParseFiles: template is nil", nil)
		}
		t.Execute(w, env)
	} else {
		http.Error(w, "File not found!", http.StatusInternalServerError)
	}
}

func main() {
	if c := os.Getenv("CFENV_BGCOLOR"); c != "" {
		bgColor = c
	}
	// We will also check to see if $PORT was set, if it was $PORT takes precedence
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Set the listening port by setting the $CFENV_PORT env variable
	if os.Getenv("CFENV_PORT") != "" {
		log.Fatal("cf-env: CFENV_PORT has been deprecated please use PORT instead!")
	}
	http.HandleFunc("/kill", killHandler)
	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
