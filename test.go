package test

import (
    "fmt"
    "log"
    "net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func main() {
    http.HandleFunc("/", helloWorldHandler)
    fmt.Println("Server is running in port 38903.")
    log.Fatal(http.ListenAndServe(":38903", nil))
}