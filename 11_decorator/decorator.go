package decorator

import "fmt"

type StrBorder struct {
	Ch byte
	Display Display
}

func (b *StrBorder) GetColumns() int {
	return 1 + b.Display.GetColumns() + 1
}

func (b *StrBorder) GetRows() int {
	return b.Display.GetRows()
}

func (b *StrBorder) GetRowText(i int) string {
	return string(b.Ch) + b.Display.GetRowText(i) + string(b.Ch)
}

func (d *StrBorder) Show() {
	rows := d.GetRows()
	for i := 0; i < rows; i++ {
		fmt.Println(d.GetRowText(i))
	}
}

type FullBorder struct {
	Display Display
}

func (b *FullBorder) GetColumns() int {
	return 1 + b.Display.GetColumns() + 1
}

func (b *FullBorder) GetRows() int {
	return 1 + b.Display.GetRows() + 1
}

func (b *FullBorder) GetRowText(i int) string {
	if i == 0 || i == b.Display.GetRows() + 1 {
		return "+" + makeLine('-', b.Display.GetColumns()) + "+"
	}
	return "|" + b.Display.GetRowText(i - 1) + "|"
}

func (d *FullBorder) Show() {
	rows := d.GetRows()
	for i := 0; i < rows; i++ {
		fmt.Println(d.GetRowText(i))
	}
}

func makeLine(ch byte, count int) string {
	res := make([]byte, count)
	for i := 0; i < count; i++ {
		res[i] = ch
	}
	return string(res)
}
