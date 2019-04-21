package common

type Output interface {
	Push([]byte) error
}
