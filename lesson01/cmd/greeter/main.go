// Command greeter prints a greeting for the name given on the command line.
//
// ENG: Usage: greeter [name]. Greets the stranger when no name is given.
//
// LAO: ທັກທາຍຕາມຊື່ທີ່ສົ່ງມາທາງ command line — ຖ້າບໍ່ມີຊື່ ຈະທັກທາຍ stranger.
package main

import (
	"fmt"
	"os"

	"example.com/lesson01/internal/greet"
)

// main reads an optional name from the arguments and prints its greeting.
//
// ENG: os.Args[0] is the program itself, so the name is os.Args[1] if present.
//
// LAO: os.Args[0] ແມ່ນໂປຣແກຣມເອງ — ຊື່ແມ່ນ os.Args[1] ຖ້າມີ.
func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greet.Hello(name))
}
