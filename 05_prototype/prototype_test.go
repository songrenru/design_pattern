package prototype

import "testing"

var manager *Manager
var productName = "productA"

func init() {
	manager = NewManager()
	manager.Set(productName, &ProductA{name: productName})
}

func TestPrototype(t *testing.T) {
	get := manager.Get(productName)
	want := get.(*ProductA)
	if want.name != productName {
		t.Error("error")
	}
}

func TestClone(t *testing.T) {
	t1 := manager.Get(productName)

	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}
