package main

import (
	"net/http"
	"log"
	//"strings"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("[GCP-HTTPS] Hello \n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

/*
func HelloServer(w http.ResponseWriter, req *http.Request){

	message := req.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "[GCP] Hello "+message

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(message))
}*/
