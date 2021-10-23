package main

/**
 * template
 */

import (
	"fmt"
	"gitee.com/yousj/golang_study/day12"
	"html/template"
	"net/http"
)

func query(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./day13/book.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	books, err := myorm.Query()
	if err != nil {
		return
	}
	terr := t.Execute(w, books)
	if terr != nil {
		fmt.Println(terr)
		return
	}
}

func main() {

	myorm.Start(query)

}
