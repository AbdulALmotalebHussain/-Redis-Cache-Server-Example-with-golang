package server

import (
    "fmt"
    "net/http"

)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
