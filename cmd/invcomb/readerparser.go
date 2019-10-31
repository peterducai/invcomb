package invcomb

import (
	"bufio"
	"fmt"
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

	file, err := os.Open(invfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") {
			fmt.Println(line)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
