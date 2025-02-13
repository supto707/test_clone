package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/html"))))

	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter , req *http.Request){
	
	temp , err := template.ParseFiles("html/index.html")

	if err != nil{
		fmt.Println(err.Error())
	}

	temp.Execute(w , nil)

}