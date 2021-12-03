package main

//https://golang.org/doc/articles/wiki/
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	if cowntdown(time.Now()) == 0 {
		fmt.Println("Merry, Christmas!")
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
http://localhost:8080/sachiko
でアクセスすると
Hi, Merry Christmas sachiko!🎅🎄✨
が表示される
*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, Merry Christmas %s!🎅🎄✨", r.URL.Path[1:])
}

func cowntdown(t time.Time) int {
	if t.Month() != time.December {
		return -1
	}
	if t.Day() < 25 {
		return 25 - t.Day()
	} else if t.Day() == 25 {
		return 0
	}
	return -1
}
