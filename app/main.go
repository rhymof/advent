package main

//https://golang.org/doc/articles/wiki/
import (
  "fmt"
  "log"
  "net/http"
)

/*
http://localhost:8080/sachiko
でアクセスすると
Hi, Merry Christmas sachiko!🎅🎄✨
が表示される
*/
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi, Merry Christmas %s!🎅🎄✨", r.URL.Path[1:])
}


func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}