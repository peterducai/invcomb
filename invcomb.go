package main

import (
	"fmt"
	"log"
	"os"

	"github.com/peterducai/invcomb/cmd/invcomb"
)

func main() {
	fmt.Println("inventory combinator")

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	invcomb.ReadFile("skdjfls")
}
