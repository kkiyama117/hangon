package external_interfaces

import (
	"net/http"
)

type output struct {
	writer http.ResponseWriter
}

func NewHttpOutput(w http.ResponseWriter) Output {
	return &output{w}
}

func (op *output) Push(data interface{}) error {
	_, err := op.writer.Write(data.([]byte))
	return err
}
