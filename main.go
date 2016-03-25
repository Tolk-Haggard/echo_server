package main

import (
	"io/ioutil"
	"net/http"

	"github.com/tolk-haggard/echo_server/writer_factory"
)

func handler(w http.ResponseWriter, r *http.Request) {
	writer := writer_factory.WriterFactory(r.UserAgent(), w)
	writer.Open()
	writer.Write("Method: %s", r.Method)
	writer.Write("Path segments: %s", r.URL.Path[1:])
	writer.Write("User Agent: %s", r.UserAgent())
	headers := r.Header
	for k := range headers {
		writer.Write("Header: %s -> %s", k, headers[k])
	}
	body, _ := ioutil.ReadAll(r.Body)
	writer.Write("Body: %s", body)
	writer.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
