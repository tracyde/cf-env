package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var bgColor = "white" // set a sane default color for background
var port = 8080       // by default listen to port 8080

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
	if p := os.Getenv("CFENV_PORT"); p != "" {
		i, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal("CFENV_PORT Conversion: ", err)
		}
		port = i
	}
	http.HandleFunc("/kill", killHandler)
	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
