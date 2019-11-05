package invcomb

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//ProcessInput process inventories
func ProcessInput(inp string) {
	// Split on comma.
	result := strings.Split(inp, ",")

	// Display all elements.
	for i := range result {
		ReadFile(result[i])
	}
}

//ReadFile to read inventories
func ReadFile(invfile string) {
	var lastGroup string

	fmt.Printf("\nreading file %s\n", invfile)

	file, err := os.Open(invfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lastGroup = "ungrouped"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		fmt.Printf("\nlast group is %s\n", lastGroup)

		// read groups
		if strings.HasSuffix(line, ":children]") {
			fmt.Printf("\nch- %s", line)
		} else if strings.HasSuffix(line, ":vars]") {
			fmt.Printf("\nv- %s", line)
		} else if strings.HasPrefix(line, "[") {
			fmt.Printf("\ng- %s", line)
			AddGroup(line)
			lastGroup = line
		} else {
			fmt.Printf("\nnode- %s", line)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func writeInventory() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	if _, err = io.WriteString(f, "hello world"); err != nil {
		fmt.Println(err)
	}

	f.Sync()
	return
}

func wr() {
	f, err := os.Create("lines")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
