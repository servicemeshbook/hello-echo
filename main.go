package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
        fmt.Fprintf(w, "The HOST IP address is: %s\n", os.Getenv("HOST_IP"))
        fmt.Fprintf(w, "The POD IP  address is: %s\n", os.Getenv("POD_IP"))
        fmt.Fprintf(w, "The HOST IP address is: %s\n", os.Getenv("POD_NAME"))
        fmt.Fprintf(w, "The HOST IP address is: %s\n", os.Getenv("POD_NAME_SPACE"))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
