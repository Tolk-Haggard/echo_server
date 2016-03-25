package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	writer := writerFactory(r.UserAgent(), w)
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

func writerFactory(userAgent string, responseWriter http.ResponseWriter) EchoWriter {
	agent := strings.ToLower(userAgent)
	if strings.Contains(agent, "curl") || strings.Contains(agent, "http") {
		return ConsoleWriter{rw: responseWriter}
	}
	return HTMLWriter{rw: responseWriter}
}

type EchoWriter interface {
	Open()
	Write(format string, data ...interface{})
	Close()
}

type HTMLWriter struct {
	rw http.ResponseWriter
}

func (h HTMLWriter) Open() {
	fmt.Fprint(h.rw, "<html><body>")
}
func (h HTMLWriter) Write(format string, data ...interface{}) {
	fmt.Fprintf(h.rw, format+" <br/>", data...)
}
func (h HTMLWriter) Close() {
	fmt.Fprint(h.rw, "</body></html>")
}

type ConsoleWriter struct {
	rw http.ResponseWriter
}

func (c ConsoleWriter) Open() {}
func (c ConsoleWriter) Write(format string, data ...interface{}) {
	fmt.Fprintf(c.rw, format+"\n", data...)
}
func (c ConsoleWriter) Close() {}
