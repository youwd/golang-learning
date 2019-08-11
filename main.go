package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8084"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	n := negroni.Classic()
	n.UseHandler(mux)
	hostString := fmt.Sprintf(":%s", port)
	n.Run(hostString)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from Go!")
}
