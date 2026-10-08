// Command cost prints the CCS2 charging cost for an energy amount in Wh.
//
// ENG: Usage: cost <energy-Wh>. Prints errors to stderr and exits with 1.
//
// LAO: ພິມລາຄາສາກ CCS2 ຕາມພະລັງງານ (Wh) — error ພິມໄປ stderr ແລະ ອອກດ້ວຍ 1.
package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/lesson02/internal/charging"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: cost <energy-Wh>")
		os.Exit(1)
	}

	eneryWh, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: energy must be a whole number of Wh:", err)
		os.Exit(1)
	}
	if eneryWh < 0 {
		fmt.Fprintln(os.Stderr, "error: energy cannot be negative")
		os.Exit(1)
	}

	cost := charging.Cost(charging.CCS2, eneryWh)
	fmt.Printf("CCS2, %d Wh: %d kip\n", eneryWh, cost)
}
