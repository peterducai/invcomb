package invcomb

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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
	fmt.Printf("%+v\n", Inv)
}

//ReadFile to read inventories
func ReadFile(invfile string) {
	var lastGroup string

	fmt.Printf("\n------------->  reading file %s\n", invfile)

	file, err := os.Open(invfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lastGroup = "ungrouped"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasSuffix(line, ":children]") {
			//CHILDREN
			var groupl = line[:len(line)-len(":children]")]
			groupl = groupl + "]"
			fmt.Printf("\ngroup %s has following children:", groupl)
			lastGroup = line
		} else if strings.HasSuffix(line, ":vars]") {
			// VARIABLES TAG
			var groupv = line[:len(line)-len(":vars]")]
			groupv = groupv + "]"
			lastGroup = line
			fmt.Printf("\ngroup %s has a vars", groupv)

		} else if strings.HasPrefix(line, "[") {
			// GROUPS
			fmt.Printf("\ng- %s", line)
			AddGroup(line)
			fileExists("groupvars/" + line + ".yml") //check if file exist!
			lastGroup = line
		} else if strings.HasPrefix(line, "#") || len(line) == 0 || line == "---" {
			// EMPTY LINES AND COMMENTS
			fmt.Println("")
		} else if strings.HasSuffix(lastGroup, ":vars]") {
			// ACTUAL VARIABLES
			fmt.Printf("\nvar- %s\n", line)
		} else if strings.HasSuffix(lastGroup, ":children]") {
			// ACTUAL VARIABLES
			fmt.Printf("\nchild- %s\n", line)
		} else {
			var vrs = strings.Split(line, " ")
			if len(vrs) == 1 {
				fmt.Printf("\nnode- %s\n", line)
			} else {
				fmt.Printf("\nnode & vars- %s\n", line)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//WriteInventory writes file
func WriteInventory(path string, single bool) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	if _, err = io.WriteString(f, "# Author: "+Inv.Author+"\n# "+Inv.Date); err != nil {
		fmt.Println(err)
	}

	f.Sync()
	return
}
