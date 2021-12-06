package main

//https://golang.org/doc/articles/wiki/
import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/Code-Hex/Neo-cowsay/v2/decoration"
	"github.com/rhymof/advent/proverbs"
)

var port = flag.Int("p", 8080, "specify port")

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Printf("Now listening on http://localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
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
		handleChristmas(w, r)
		return
	}
	handleCowntdown(w, r)
}

func handleChristmas(w http.ResponseWriter, r *http.Request) {
	name := getNameFromRequest(r)
	msg := fmt.Sprintf("Hi, Merry Christmas %s!🎅🎄✨", name)
	msg = decorateFunc(msg)
	fmt.Fprintf(decorateWriterFunc(w), msg)
}

func handleCowntdown(w http.ResponseWriter, r *http.Request) {
	name := getNameFromRequest(r)
	msg := fmt.Sprintf("Hi, %s! It's %d days until Christmas!🦌✨", name, cowntdown(nowFunc()))
	fmt.Fprintf((w), msg)
	proverb := proverbs.FromDate(cowntdown(nowFunc()))
	fmt.Fprintf((w), proverb)
}

func getNameFromRequest(r *http.Request) string {
	return r.URL.Path[1:]
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
