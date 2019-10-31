package invcomb

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//jumper ansible_port=5555 ansible_host=192.0.2.50
//[databases]
//db-[a:f].example.com

//[targets]
//
//localhost              ansible_connection=local
//other1.example.com     ansible_connection=ssh        ansible_user=mpdehaan
//other2.example.com     ansible_connection=ssh        ansible_user=mdehaan

// [atlanta]
// host1
// host2

// [atlanta:vars]
// ntp_server=ntp.atlanta.example.com
// proxy=proxy.atlanta.example.com

// [southeast:vars]
// some_server=foo.southeast.example.com
// halon_system_timeout=30
// self_destruct_countdown=60
// escape_pods=2

//GroupTree represents tree of groups, subgroups and end nodes (hosts) with variables
type GroupTree struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

//MasterGroup defines master group (All)
type MasterGroup struct {
	Grouped   GroupTree
	Ungrouped GroupTree
}

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
