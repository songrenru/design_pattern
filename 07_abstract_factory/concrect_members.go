package abstract_factory

import "fmt"

type RDBMainDAO struct {}

func (r *RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main save")
}

type RDBDetailDAO struct {}

func (r *RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

type RDBFactory struct {}

func (f *RDBFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (f *RDBFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct {}

func (r *XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

type XMLDetailDAO struct {}

func (r *XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

type XMLFactory struct {}

func (f *XMLFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (f *XMLFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}