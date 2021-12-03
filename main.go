package main

//https://golang.org/doc/articles/wiki/
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

/*
http://localhost:8080/sachiko
でアクセスすると
Hi, Merry Christmas sachiko!🎅🎄✨
が表示される
*/
var nowFunc = time.Now

func handler(w http.ResponseWriter, r *http.Request) {
	if isChristmas(nowFunc()) {
		fmt.Fprintf(w, "Hi, Merry Christmas %s!🎅🎄✨", r.URL.Path[1:])
	}
}

func isChristmas(t time.Time) bool {
	return cowntdown(t) == 0
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
