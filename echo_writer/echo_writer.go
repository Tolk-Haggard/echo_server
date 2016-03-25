package echo_writer

import (
	"fmt"
	"net/http"
)

type EchoWriter interface {
	Open()
	Write(format string, data ...interface{})
	Close()
}

type HTMLWriter struct {
	RW http.ResponseWriter
}

func (h HTMLWriter) Open() {
	fmt.Fprint(h.RW, "<html><body>")
}
func (h HTMLWriter) Write(format string, data ...interface{}) {
	fmt.Fprintf(h.RW, format+" <br/>", data...)
}
func (h HTMLWriter) Close() {
	fmt.Fprint(h.RW, "</body></html>")
}

type ConsoleWriter struct {
	RW http.ResponseWriter
}

func (c ConsoleWriter) Open() {}
func (c ConsoleWriter) Write(format string, data ...interface{}) {
	fmt.Fprintf(c.RW, format+"\n", data...)
}
func (c ConsoleWriter) Close() {}
