package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	b,err := ioutil.ReadFile("./test/main.go")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v",err)))
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/test/", test)
	http.ListenAndServe(":5051", nil)
}
