package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/f1/", f1)
	http.HandleFunc("/v1/f2/", f2)
	http.ListenAndServe("127.0.0.1:8085", nil)
}

func f1(w http.ResponseWriter, r *http.Request) {
	// str := "<a href='https://www.baidu.com' target='_blank'>hello go httpServer</a>"
	// w.Write([]byte(str))
	bin, err := ioutil.ReadFile("./index.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.Write(bin)
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	w.Write([]byte("OK"))
}
