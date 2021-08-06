package decorator

import "fmt"

type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(i int) string
	Show()
}

type StrDisplay struct {
	Str string
}

func (d *StrDisplay) GetColumns() int {
	return len(d.Str)
}

func (d *StrDisplay) GetRows() int {
	return 1
}

func (d *StrDisplay) GetRowText(i int) string {
	return d.Str
}

func (d *StrDisplay) Show() {
	rows := d.GetRows()
	for i := 0; i < rows; i++ {
		fmt.Println(d.GetRowText(i))
	}
}
