package invcomb

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
