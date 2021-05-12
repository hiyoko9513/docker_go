package main
import (
	"fmt"
	"log"
	"net/http"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() 
	fmt.Println("----Http.Request----")
	fmt.Println("----line----")
	fmt.Println(r.Method, r.RequestURI, r.Proto)
	fmt.Println("----header----")
	fmt.Println(r.Header)
	fmt.Println("----body----")
	fmt.Println(r.Body)
	fmt.Println("----end----")
	fmt.Fprintf(w, "hello go server!")
}
func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Listen in 8080.")
	if err != nil {
		log.Fatal("Listen and server:", err)
	}
}
