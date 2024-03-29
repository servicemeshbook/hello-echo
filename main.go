package main

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "log"
    "os"
    "time"
)

func main() {
    log.SetFlags(log.Ldate | log.Lmicroseconds)
    log.Println("main started")

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        log.Println("In /health")
        fmt.Fprintf(w, "OK")
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("In /")
        fmt.Fprintf(w, "===============================================\n")
        fmt.Fprintf(w, "Request time   : %s\n", time.Unix(0, time.Now().UnixNano())) 
        fmt.Fprintf(w, "Requested path : %s\n", r.URL.Path)
        fmt.Fprintf(w, "Host IP        : %s\n", os.Getenv("HOST_IP"))
        fmt.Fprintf(w, "Pod IP         : %s\n", os.Getenv("POD_IP"))
        fmt.Fprintf(w, "Pod Name       : %s\n", os.Getenv("POD_NAME"))
        fmt.Fprintf(w, "Pod Namespace  : %s\n", os.Getenv("POD_NAMESPACE"))
        fmt.Fprintf(w, "Host           : %s\n", r.Host)
        fmt.Fprintf(w, "RemoteAddr     : %s\n", r.RemoteAddr)
        fmt.Fprintf(w, "===============================================\n")

        upstreamService := os.Getenv("UPSTREAM_SERVICE")
        if (len(upstreamService) > 0) {
          log.Printf("Call to upstream service %s\n", upstreamService)
          response, err := http.Get(upstreamService)
          if err != nil {
            fmt.Printf("The response from %s did not come through %s\n", upstreamService, err)
          } else {
            data, _ := ioutil.ReadAll(response.Body)
            log.Printf("Response starts from %s\n", upstreamService)
            fmt.Fprintf(w,string(data))
            log.Printf("Response ends from %s\n", upstreamService)
          }
        } else {
          log.Println("No upstream service defined")
        }
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
