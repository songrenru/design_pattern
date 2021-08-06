package facade

import "fmt"

// 门面入口
type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// 具体门面
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// 门面背后的逻辑

//AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

//BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
