package main

import (
    "bytes"
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "os"
    "time"
    "io/ioutil"
)

func main() {

    log.SetFlags(log.Ldate | log.Lmicroseconds)
    log.Println("main started")

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        log.Printls("In /health")
        fmt.Fprintf(w, "OK")
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "===============================================\n")
        fmt.Fprintf(w, "Request time   : %s\n", time.Unix(0, time.Now().UnixNano())) 
        fmt.Fprintf(w, "Requested path : %s\n", r.URL.Path)
        fmt.Fprintf(w, "===============================================\n")

        //for k, v := range r.Header {
        //   fmt.Fprintf(w, "%q: %q\n", k, v)
        //}

        jsonData := map[string]string{"POD_IP": os.Getenv("POD_IP"), 
                                      "POD_NAME": os.Getenv("POD_NAME"),
                                      "POD_NAMESPACE": os.Getenv("POD_NAMESPACE"),
                                      "HOST_IP": os.Getenv("HOST_IP"),
                                      "HOST": r.Host,
                                      "REMOTE_ADDR": r.RemoteAddr}
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
