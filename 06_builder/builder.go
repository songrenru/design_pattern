package builder

type Builder interface {
	part1()
	part2()
	part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) part1()  {
	b.result += "1"
}

func (b *Builder1) part2()  {
	b.result += "2"
}

func (b *Builder1) part3()  {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) part1()  {
	b.result += 1
}

func (b *Builder2) part2()  {
	b.result += 2
}

func (b *Builder2) part3()  {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}