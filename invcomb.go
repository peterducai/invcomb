package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	fmt.Println("inventory combinator")

	// invInput := flag.String("input", "define_input_inventories", "inventories to combine")
	// invOutput := flag.String("output", "examples/generated_inventory.yml", "inventory to create from input inventories")
	// singleFile := flag.Bool("singlefile", true, "create single inventory file with all vars included.. dont generate host_vars and group_vars")

	inputfile := flag.String("inputfile", "define_input_inventories", "inventories to combine")
	inputfolder := flag.String("inputfolder", "define_input_inventories", "inventories to combine")
	outputfile := flag.String("outputfile", "define_input_inventories", "inventories to combine")
	outputfolder := flag.String("outputfolder", "define_input_inventories", "inventories to combine")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\n Using %s,%s,%s,%s in %s\n", *inputfile, *inputfolder, *outputfile, *outputfolder, dir)

	//PrintMemUsage()

	// invcomb.InitInventory(*invOutput)

	// invcomb.ProcessInput(*invInput)

	// invcomb.WriteInventory(*invOutput, *singleFile)

	// fmt.Printf("\n\ngoing to create %s in %s\n", *invOutput, dir)

	// runtime.GC()
	//PrintMemUsage()
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
