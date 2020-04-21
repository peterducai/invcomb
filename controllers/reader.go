package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/peterducai/invcomb/models"
)

//ReadDir reads directory group_vars and host_vars
func ReadDir(rootdir string) {

}

//ReadFile reads inventory file
func ReadFile(invfile string) {

	fmt.Printf("\n------------->  reading file %s\n", invfile)

	file, err := os.Open(invfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasSuffix(line, ":children]") {
			//CHILDREN
			var groupl = line[:len(line)-len(":children]")]
			groupl = groupl + "]"
			//fmt.Printf("\ngroup %s has following children:", groupl)
			lastGroup = line
			childrenpart = true
		} else if strings.HasSuffix(line, ":vars]") {
			// VARIABLES TAG
			var groupv = line[:len(line)-len(":vars]")]
			groupv = groupv + "]"
			lastGroup = line
			//fmt.Printf("\ngroup %s has a vars", groupv)

		} else if strings.HasPrefix(line, "[") {
			// GROUPS
			//fmt.Printf("\ng- %s", line)
			models.AddGroup(line)
			fileExists("groupvars/" + line + ".yml") //check if file exist!
			lastGroup = line
		} else if strings.HasPrefix(line, "#") || len(line) == 0 || line == "---" {
			// EMPTY LINES AND COMMENTS
			//fmt.Println("")
		} else if strings.HasSuffix(lastGroup, ":vars]") {
			// ACTUAL VARIABLES
			//fmt.Printf("\nvar- %s\n", line)
		} else if strings.HasSuffix(lastGroup, ":children]") {
			// ACTUAL CHILDREN
			//fmt.Printf("\nchild- %s\n", line)
		} else {
			var vrs = strings.Split(line, " ")
			if len(vrs) == 1 {
				//fmt.Printf("\nnode- %s\n", line)
			} else {
				//fmt.Printf("\nnode & vars- %s\n", line)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
