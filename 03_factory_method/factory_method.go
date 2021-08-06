package factory_method

// go的工厂方法模式，模板方法独立了出来，见test的compute()

// 抽象产品,是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 抽象工厂
type OperatorFactory interface {
	Create() Operator
}

//OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

//PlusOperator 的实现
type PlusOperator struct {
	*OperatorBase
}

func (p PlusOperator) Result() int {
	return p.a + p.b
}

//PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct {}

func (f PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//MinusOperator 的实现
type MinusOperator struct {
	*OperatorBase
}

func (p MinusOperator) Result() int {
	return p.a - p.b
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct {}

func (f MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

