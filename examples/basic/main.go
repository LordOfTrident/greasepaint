package main

import (
	"fmt"

	gp "github.com/LordOfTrident/greasepaint"
)

func main() {
	fmt.Println(gp.Center("Hello, world"))
	fmt.Println(gp.Center(gp.LightBlueFg("This is a nice blue text.")))
	fmt.Println(gp.Chain(gp.LightRedFg, gp.Center)("This is a nice red text!"))

	myStyle := gp.NewStyle().
		Bold       (true).
		Italic     (true).
		Underline  (true).
		Fg         (gp.Black).
		Bg         (gp.White).
		Padding    (1).
		BorderLeft ('>').
		BorderRight('<')

	fmt.Println(gp.Center(myStyle.Apply("This is a fancy styled text")))

	fmt.Println()
	fmt.Println(gp.Left(gp.Gap(4) + "Im aligned to the left!"))
	fmt.Println(gp.Right("Im aligned to the right!" + gp.Gap(4)))
	fmt.Println()

	styleGrid := gp.NewStyle().
		BgHex  ("eceff2").
		Padding(4).
		GridGap(1).
		Apply

	styleElem := gp.NewStyle().
		Bold   (true).
		BgHex  ("2c94e9").
		FgHex  ("fff").
		Padding(1).
		Apply

	fmt.Println(gp.Center(styleGrid(styleElem("Apple"), styleElem("Orange"), styleElem("Banana"))))
}
