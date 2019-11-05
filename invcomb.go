package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

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

	PrintMemUsage()

	invcomb.InitInventory(*invOutput)

	invcomb.ProcessInput(*invInput)

	fmt.Printf("\n\ngoing to create %s in %s\n", *invOutput, dir)

	// runtime.GC()
	PrintMemUsage()
}

//PrintMemUsage show memory usage
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
