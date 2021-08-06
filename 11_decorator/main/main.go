package main

import decorator "github.com/songrenru/design_attern/11_decorator"

func main() {
	d := &decorator.StrDisplay{Str: "hello, world"}
	b1 := &decorator.StrBorder{Ch: '#', Display: d}
	b2 := &decorator.FullBorder{Display: b1}
	d.Show()
	b1.Show()
	b2.Show()
}
