package simple_factory

import "testing"

func TestHi(t *testing.T) {
	api := NewAPI(TypeHi)
	get := api.Say("eason")
	want := "hi, eason"
	if get != want {
		t.Fatal("TestHi fail")
	}
}

func TestHello(t *testing.T) {
	api := NewAPI(TypeHello)
	get := api.Say("eason")
	want := "hello, eason"
	if get != want {
		t.Fatal("TestHello fail")
	}
}
