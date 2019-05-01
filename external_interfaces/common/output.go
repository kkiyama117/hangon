package common

type Output interface {
	Push(interface{}) error
}
