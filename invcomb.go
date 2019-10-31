package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/peterducai/invcomb/cmd/invcomb"
)

func main() {
	fmt.Println("inventory combinator")

	invInput := flag.String("input", "define_input_inventories", "inventories to combine")
	invOutput := flag.String("output", "examples/generated_inventory.yml", "inventory to create from input inventories")
	//boolPtr := flag.Bool("fork", false, "a bool")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	invcomb.ProcessInput(*invInput)

	fmt.Printf("\n\ngoing to create %s in %s\n", *invOutput, dir)
}
