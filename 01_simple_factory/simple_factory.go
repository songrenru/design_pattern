package simple_factory

import "fmt"

const (
	TypeHi int = iota
	TypeHello
)

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	switch t {
	case TypeHi:
		return &hiAPI{}
	case TypeHello:
		return &helloAPI{}
	}

	return nil
}

type hiAPI struct {}

func (h hiAPI) Say(name string) string {
	return fmt.Sprintf("hi, %s", name)
}

type helloAPI struct {}

func (h helloAPI) Say(name string) string {
	return fmt.Sprintf("hello, %s", name)
}
