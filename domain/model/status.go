package model

import "errors"

type Status struct {
	id       int
	Text       string
	StatusType statusType
}

type statusType string

const (
	SuccessStatus statusType = "success"
	ErrorStatus   statusType = "error"
)

// return error from standard error package
func (status *Status) Error() error {
	if status.StatusType == ErrorStatus {
		return errors.New(status.Text)
	}
	return nil
}
