package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/peterducai/invcomb/controllers"
	"github.com/peterducai/invcomb/models"
)

func main() {
	fmt.Println("inventory combinator")

	// invInput := flag.String("input", "define_input_inventories", "inventories to combine")
	// invOutput := flag.String("output", "examples/generated_inventory.yml", "inventory to create from input inventories")
	// singleFile := flag.Bool("singlefile", true, "create single inventory file with all vars included.. dont generate host_vars and group_vars")

	inputfile := flag.String("inputfile", "", "inventory to read")
	inputfolder := flag.String("inputfolder", "", "folder to read")
	outputfile := flag.String("outputfile", "generated.yml", "outputfile")
	outputfolder := flag.String("outputfolder", "", "folder to generate")
	flag.Parse()

	if *inputfile == "" {
		if *inputfolder == "" {
			fmt.Println("missing input")
			os.Exit(1)
		}

	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\n Using %s,%s,%s,%s in %s\n", *inputfile, *inputfolder, *outputfile, *outputfolder, dir)

	//PrintMemUsage()

	models.InitInventory(*outputfile)

	controllers.ProcessInput(*inputfile)

	if *outputfile != "" {
		controllers.WriteSingleInventoryFile(*outputfile, true)
	}

	//fmt.Printf("\n\ngoing to create %s in %s\n", *outputfile, dir)

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
