package static_server

// serve http
import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Based Go Microservice")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	dir := http.Dir("./static")
	port := 8080
	serveString := ":" + string(rune(port))
	fileServer := http.FileServer(dir)
	http.Handle("/", fileServer)
	http.HandleFunc("/test", handler)

	err := http.ListenAndServe(serveString, nil)
	if err != nil {
		panic(err)
	}
}
