package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pumpkin/external_interfaces/common"
)

type output struct {
	writer http.ResponseWriter
}

func NewViewOutput(w http.ResponseWriter) common.Output {
	return &output{w}
}

func (op *output) Push(data []byte) error {
	fmt.Printf("%v\n", json.NewEncoder(op.writer).Encode(data))
	return nil
}
