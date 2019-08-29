package main

import (
    "bytes"
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "os"
    "io/ioutil"
)

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

        jsonData := map[string]string{"POD_IP": os.Getenv("POD_IP") , "POD_NAME": os.Getenv("POD_NAME") }
        jsonValue, _ := json.Marshal(jsonData)
        response, err := http.Post("http://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
        if err != nil {
          fmt.Printf("The response from http://httpbin.org/post did not come through %s\n", err)
        } else {
          data, _ := ioutil.ReadAll(response.Body)
          fmt.Fprintf(w, "Response starts from http://httpbin.org/post\n")
          fmt.Fprintf(w,string(data))
          fmt.Fprintf(w, "Response ends from http://httpbin.org/post\n")
        } 
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
