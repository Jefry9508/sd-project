package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "Yo soy %s\n", hostname)
	fmt.Fprintf(w, "Yo soy %s\n", hostname)
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}
