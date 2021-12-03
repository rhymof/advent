package main

//https://golang.org/doc/articles/wiki/
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/Code-Hex/Neo-cowsay/v2/decoration"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

// https://github.com/Code-Hex/Neo-cowsay/tree/master/cows
const recommendType = "dragon-and-cow"

var (
	nowFunc            = time.Now
	decorateFunc       = decorate
	decorateWriterFunc = decorateWriter
)

func decorate(msg string) string {
	mow, _ := cowsay.Say(msg, cowsay.Type(recommendType))
	return mow
}

func decorateWriter(w io.Writer) io.Writer {
	return decoration.NewWriter(
		w,
		decoration.WithBold(),
		decoration.WithRainbow(),
	)
}

/*
http://localhost:8080/sachiko
でアクセスすると
Hi, Merry Christmas sachiko!🎅🎄✨
が表示される
*/
func handler(w http.ResponseWriter, r *http.Request) {
	if isChristmas(nowFunc()) {
		msg := fmt.Sprintf("Hi, Merry Christmas %s!🎅🎄✨", r.URL.Path[1:])
		msg = decorateFunc(msg)
		fmt.Fprintf(decorateWriterFunc(w), msg)
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
