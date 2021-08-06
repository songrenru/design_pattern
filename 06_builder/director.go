package builder

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.part1()
	d.builder.part2()
	d.builder.part3()
}

