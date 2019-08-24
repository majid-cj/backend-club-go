package main
 
import (
    "fmt"
    "log"
    "net/http"
)
func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/test" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
 
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, "why-it-is-greate/web/get-post/form.html")
    case "POST":
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
        fmt.Fprintf(w, "Name = %s\n", r.FormValue("name"))
        fmt.Fprintf(w, "Address = %s\n", r.FormValue("address"))
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}
 
func main() {
    http.HandleFunc("/test", hello)
 
    fmt.Printf("Starting server for testing HTTP GET/POST...\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}