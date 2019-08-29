package main

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "log"
    "os"
)

func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "===============================================\n")
        fmt.Fprintf(w, "Requested path    is: %s\n", r.URL.Path)
        fmt.Fprintf(w, "The HOST IP       is: %s\n", os.Getenv("HOST_IP"))
        fmt.Fprintf(w, "The POD IP        is: %s\n", os.Getenv("POD_IP"))
        fmt.Fprintf(w, "The POD NAME      is: %s\n", os.Getenv("POD_NAME"))
        fmt.Fprintf(w, "The POD NAMESPACE is: %s\n\n", os.Getenv("POD_NAMESPACE"))
        fmt.Fprintf(w, "===============================================\n")

        for k, v := range r.Header {
           fmt.Fprintf(w, "%q: %q\n", k, v)
        }

        fmt.Fprintf(w, "===============================================\n")
        fmt.Fprintf(w, "Host       = %q\n", r.Host)
        fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
        fmt.Fprintf(w, "===============================================\n")
        upstreamService := getEnv("UPSTREAM_SERVICE", "http://httpbin.org/headers") 
        response, err := http.Get(upstreamService)
        if err != nil {
          fmt.Printf("The response from %s did not come through %s\n", upstreamService, err)
        } else {
          data, _ := ioutil.ReadAll(response.Body)
          fmt.Fprintf(w, "Response starts from %s\n", upstreamService)
          fmt.Fprintf(w,string(data))
          fmt.Fprintf(w, "Response ends from %s\n", upstreamService)
        }
    })


    log.Fatal(http.ListenAndServe(":8080", nil))
}
