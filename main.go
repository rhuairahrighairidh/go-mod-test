package main

import (
	"fmt"

	depa "github.com/rhuairahrighairidh/go-mod-test-depA"
	depb "github.com/rhuairahrighairidh/go-mod-test-depB"
)

func main() {
	fmt.Println("This is a go mod test")
	fmt.Println("It uses two dependencies: A and B")
	fmt.Printf("Dep A has version: %s\n", depa.GetVersion())
	fmt.Printf("Dep B has version: %s\n", depb.GetVersion())
}
