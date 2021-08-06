package main

import bridge "github.com/songrenru/design_attern/08_bridge"

func main() {
	impl := &bridge.StrDisplayImpl{Str: "bridge pattern"}
	display := bridge.Display{
		Impl: impl,
	}

	display.Show()

	countDisplay := bridge.CountDisplay{
		Display: &display,
		Count: 5,
	}
	countDisplay.Show()
}
