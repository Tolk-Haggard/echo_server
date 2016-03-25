package writer_factory

import (
	"net/http"
	"strings"

	"github.com/tolk-haggard/echo_server/echo_writer"
)

func WriterFactory(userAgent string, responseWriter http.ResponseWriter) echo_writer.EchoWriter {
	agent := strings.ToLower(userAgent)
	if strings.Contains(agent, "curl") || strings.Contains(agent, "http") {
		return echo_writer.ConsoleWriter{RW: responseWriter}
	}
	return echo_writer.HTMLWriter{RW: responseWriter}
}
