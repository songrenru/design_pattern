package bridge

import "fmt"

// 实现层
type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}

// 功能层
type Display struct {
	Impl DisplayImpl // 委托，弱关联，桥接起功能与实现
}

func (d *Display) Open() {
	d.Impl.rawOpen()
}

func (d *Display) Print() {
	d.Impl.rawPrint()
}

func (d *Display) Close() {
	d.Impl.rawClose()
}

func (d *Display) Show() {
	d.Open()
	d.Print()
	d.Close()
}

// 实现层实现
type StrDisplayImpl struct {
	Str string
}

func (i *StrDisplayImpl) rawOpen() {
	fmt.Println("-------")
}

func (i *StrDisplayImpl) rawPrint() {
	fmt.Println(i.Str)
}

func (i *StrDisplayImpl) rawClose() {
	fmt.Println("-------")
}

// 功能层扩展
type CountDisplay struct {
	*Display
	Count int
}

func (d *CountDisplay) Show() {
	d.Open()
	for i := 0; i < d.Count; i++ {
		d.Print()
	}
	d.Close()
}


