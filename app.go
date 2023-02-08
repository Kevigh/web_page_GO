package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Country string
	Age     int
}

func Index(rw http.ResponseWriter, r *http.Request) {
	template, error := template.ParseFiles("templates/index.html")

	user := User{"Kevin Giraldo Hincapié", "Colombia", 20}

	if error != nil {
		panic(error)
	} else {
		template.Execute(rw, user)
	}
}

func main() {

	http.HandleFunc("/", Index)

	fmt.Println("The server is run localhost:3000")
	fmt.Println("Run server: http://localhost:3000")

	http.ListenAndServe("localhost:3000", nil)

}
