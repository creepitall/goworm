package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/creepitall/goworm/internal/app"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	runner := app.New()
	go runner.Run()

	http.Header.Add(http.Header{}, "Content-Type", "text/plain")
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/main.js", serveJs)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		app.ServeWs(runner, w, r)
	})
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./web/index.html")

}

func serveJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/main.js")
}
