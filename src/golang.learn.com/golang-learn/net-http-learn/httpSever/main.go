package main

import "net/http"

func main() {
	http.HandleFunc("/v1/posts/Go/req1/", f1)
	http.ListenAndServe("127.0.0.1:8085", nil)
}

func f1(w http.ResponseWriter, r *http.Request) {
	str := "<a href='https://www.baidu.com' target='_blank'>hello go httpServer</a>"
	w.Write([]byte(str))
}
