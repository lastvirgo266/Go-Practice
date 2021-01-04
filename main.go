package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServerHTTP(writer http.ResponseWriter, reader *http.Request){
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(writer, nil)
}

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(
			`
			<html>
				<head>
				<title>Chat</title>
				</head>
			</html>
			`))
	})

	if err := http.ListenAndServe(":8080", nil);
	err != nil{
		log.Fatal("ListenAndServer : ", err)
	}

}
