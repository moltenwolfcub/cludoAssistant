package main

import (
	"fmt"

	"github.com/moltenwolfcub/cludoAssistant/cludo"
)

func main() {
	who := cludo.NewQuestion(
		cludo.NewOption("Green"),
		cludo.NewOption("Mustard"),
		cludo.NewOption("Peacock"),
		cludo.NewOption("Plum"),
		cludo.NewOption("Scarlet"),
		cludo.NewOption("White"),
	)
	fmt.Println(who)
}
