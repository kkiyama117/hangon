package model

import "errors"

type Status struct {
	id         int
	StatusText statusText
}

type statusText string

const (
	SuccessStatusType statusText = "success"
	ErrorStatusType   statusText = "error"
)

// return error from standard error package
func (status *Status) Error() error {
	if status.StatusText != SuccessStatusType {
		return nil
	}
	return errors.New(string(status.StatusText))
}
