package prototype

//Cloneable 是原型对象需要实现的接口
type Product interface {
	Clone() Product
}

// 原型对象应该分配在不同package，这里简单处理了
type ProductA struct {
	name string
}

func (a *ProductA) Clone() Product {
	newA := *a
	return &newA
}

type ProductB struct {
	name string
}

func (b *ProductB) Clone() Product {
	newB := *b
	return &newB
}
