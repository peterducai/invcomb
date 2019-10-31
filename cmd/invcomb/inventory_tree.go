package invcomb

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//xProcessInput process inventories
func xProcessInput(inp string) {
	// Split on comma.
	result := strings.Split(inp, ",")

	// Display all elements.
	for i := range result {
		ReadFile(result[i])
	}
}

//xReadFile to read inventories
func xReadFile(invfile string) {
	fmt.Printf("\nreading file %s\n", invfile)

	file, err := os.Open(invfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") {
			fmt.Printf("\n- %s", line)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
