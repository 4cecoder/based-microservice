package static_server

// serve http
import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Based Go Microservice")
}
func main() {
    port := 8080
    serveString := ":" + string(rune(port))
    http.HandleFunc("/", handler)

    err := http.ListenAndServe(serveString, nil)
    if err != nil {
        panic(err)
    }
}
