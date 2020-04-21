package controllers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/peterducai/invcomb/models"
)

var lastGroup = "ungrouped"
var childrenpart = false
var variablespart = false

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//ProcessInput process inventories
func ProcessInput(inp string) {
	// Split on comma.
	result := strings.Split(inp, ",")

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var allf = dir
	allf = allf + "/group_vars/all"
	if fileExists(allf) {
		ReadFile(allf)
	} else {
		fmt.Printf("\n%s NOT FOUND\n", allf)
	}

	// Display all elements.
	for i := range result {
		ReadFile(result[i])
	}

	fmt.Println("===END==============================================")
	fmt.Printf("%+v\n", models.Inv)
}
