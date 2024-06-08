package main

import (
	"fmt"
	"os"

	forge "github.com/guillaumemichel/peerid-forge"
)

func main() {
	if len(os.Args) != 2 || os.Args[1][0] == '-' {
		fmt.Println("  Usage: ./forge <base58 peerid suffix>\n  Example: ./forge ooooPEER\n          > 1EooooPEER")
		os.Exit(0)
	}
	forged, err := forge.ForgePeerID(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	fmt.Println(forged)
}
