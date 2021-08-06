package prototype

type Manager struct {
	prototypes map[string]Product
}

func NewManager() *Manager {
	return &Manager{prototypes: map[string]Product{}}
}

func (m *Manager) Set(name string, prototype Product) {
	m.prototypes[name] = prototype
}

func (m *Manager) Get(name string) Product {
	return m.prototypes[name]
}



