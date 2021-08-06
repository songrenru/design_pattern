package iterator

// 需要迭代集合的定义
type Aggregate interface {
	Iterator() Iterator
}

// 迭代器定义
type Iterator interface {
	HasNext() bool
	Next() interface{}
}


