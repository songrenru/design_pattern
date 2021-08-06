package main

import singleton "github.com/songrenru/design_attern/04_singleton"

func main() {
	s := singleton.GetInstance()
	s.Do()
}
