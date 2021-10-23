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
	db, err := dbtool.ConnMysql()
	if err != nil {
		fmt.Println(err)
		return
	}
	books := dbtool.Query(db)
	tErr := t.Execute(w, books)
	if tErr != nil {
		fmt.Println(tErr)
		return
	}
}

func start(handler func(http.ResponseWriter, *http.Request)) error {
	http.HandleFunc("/book", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	start(query)
}
