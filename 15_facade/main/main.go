package main

import (
	"fmt"
	facade "github.com/songrenru/design_attern/15_facade"
)

func main() {
	api := facade.NewAPI()
	fmt.Println(api.Test())
}
