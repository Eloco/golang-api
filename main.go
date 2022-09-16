//package main
//
//import (
//	//"errors"
//	"fmt"
//	"io"
//	"net/http"
//	//"os"
//)
//
//func getRoot(w http.ResponseWriter, r *http.Request) {
//	fmt.Printf("got / request\n")
//	io.WriteString(w, "This is my website!\n")
//}
//func getHello(w http.ResponseWriter, r *http.Request) {
//	fmt.Printf("got /hello request\n")
//	io.WriteString(w, "Hello, HTTP!\n")
//}
//
//func main() {
//	http.HandleFunc("/", getRoot)
//	http.HandleFunc("/hello", getHello)
//
//	err := http.ListenAndServe(":3333", nil)
//	fmt.Println(err)
//}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is golang api server!\n")
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":

		// gain client post raw body
		//body, err := ioutil.ReadAll(r.Body) // if body is big, ioutil.ReadAll not work well
		//if err != nil {
		//	fmt.Println("read body error：", err)
		//	return
		//}
		//fmt.Println(string(body))

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		//fmt.Println("path", r.URL.Path)
		//fmt.Println("scheme", r.URL.Scheme)

		//fmt.Println(r.Form)
		//for k, v := range r.Form {
		//	fmt.Println("key:", k)
		//	fmt.Println("val:", strings.Join(v, ""))
		//}

		run := strings.Join(r.Form["run"], "")
		fmt.Println(run)

		// give client JSON data
		result := make(map[string]interface{})
		result["param1"] = "result1"
		result["param2"] = "result2"
		w.Header().Set("Content-Type", "application/json")
		json, _ := json.Marshal(result)
		w.Write(json)
	default:
		io.WriteString(w, "we only support POST method!\n")
	}
}
func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/post", ApiHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
