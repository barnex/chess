package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var flagAddr = flag.String("http", "localhost:26355", "HTTP address")

func main() {
	http.Handle("/play", http.StripPrefix("/play", http.HandlerFunc(handlePlay)))
	fmt.Println("visit http://" + *flagAddr + "/play")
	log.Fatal(http.ListenAndServe(*flagAddr, nil))
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "chess?")
}
