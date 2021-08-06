package abstract_factory

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDAO interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}
