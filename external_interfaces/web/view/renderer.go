package view

import (
	"net/http"

	"pumpkin/external_interfaces/common"
)

type output struct {
	writer http.ResponseWriter
}

func NewViewOutput(w http.ResponseWriter) common.Output {
	return &output{w}
}

func (op *output) Push(data interface{}) error {
	_, err := op.writer.Write(data.([]byte))
	return err
}
