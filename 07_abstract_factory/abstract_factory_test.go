package abstract_factory

import "testing"

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestAbsFactory(t *testing.T) {
	t.Run("rdb", func(t *testing.T) {
		var f DAOFactory = &RDBFactory{}
		getMainAndDetail(f)
	})

	t.Run("rdb", func(t *testing.T) {
		var f DAOFactory = &XMLFactory{}
		getMainAndDetail(f)
	})
}
