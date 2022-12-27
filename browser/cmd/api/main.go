package main

import (
    "io"

    "log"
    "net/http"
    _ "net/http/pprof"
)

type Config struct {
    buf []byte
}

func main() {
    app := Config{}

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Boo! \n")
    })

    http.HandleFunc("/picture", app.GetPicture)

    go func() {
        log.Println(http.ListenAndServe(":6060", nil))
    }()

    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal(err)
    }
}