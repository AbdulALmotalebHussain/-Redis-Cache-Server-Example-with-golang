package main

import (
	"fmt"

     "caching/internal/cache"
     "caching/internal/server"
     "net/http"
)

func main() {

    cache.InitRedis("localhost:6379", "")

    http.HandleFunc("/", server.HomeHandler)
    fmt.Println("Listening on port 8080")

    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Println("error listening and serving: ")

    }

    fmt.Println("Hello, World!")

}
