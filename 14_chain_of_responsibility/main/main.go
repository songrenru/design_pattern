package main

import chain_of_responsibility "github.com/songrenru/design_attern/14_chain_of_responsibility"

func main() {
	limit := chain_of_responsibility.NewLimitSupport(100)
	odd := chain_of_responsibility.NewOddSupport()
	sepecial := chain_of_responsibility.NewSepecialSupport(298)
	limit.SetNext(odd).SetNext(sepecial)

	t := chain_of_responsibility.NewTrouble(0)
	for i := 1; i < 300; i++ {
		t.Num = i
		limit.Resolve(t)
	}
}
