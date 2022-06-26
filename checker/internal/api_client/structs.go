package api_client

import "fmt"

var HTTP503Error = fmt.Errorf("HTTP 503")

type CompilerError struct {
	Error string `json:"error"`
	Line  int    `json:"line"`
}
