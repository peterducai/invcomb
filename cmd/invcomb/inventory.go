package invcomb

import (
	"fmt"
	"os/user"
	"time"
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

//Variable is var with = value
type Variable struct {
	vr string
}

//Node or host
type Node struct {
	Name      string
	Variables []Variable
}

//Group in inventory
type Group struct {
	Name      string
	Nodes     []Node
	Children  []string
	Variables []Variable
}

//Inventory represents tree of groups, subgroups and end nodes (hosts) with variables
type Inventory struct {
	Name      string
	Date      string
	Author    string
	Groups    []Group
	Nodes     []Node
	Variables []Variable
}

//Inv inventory
var Inv Inventory

//InitInventory initialize inventory
func InitInventory(nm string) {
	fmt.Printf("INIT OF INVENTORY %s ------------------\n\n", nm)
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	Inv.Date = time.Now().Format(time.RFC850)
	Inv.Name = nm
	Inv.Author = user.Name

	fmt.Printf("%+v\n", Inv)
}

//AddGroup add group to inventory  TODO: fix
func AddGroup(g string) {
	// fmt.Println("adding group --------------- ")
	for _, s := range Inv.Groups {
		fmt.Printf("> %s", s)
		if s.Name != g {
			var gr = new(Group)
			gr.Name = g
			Inv.Groups = append(Inv.Groups, *gr)
		}
	}

}

//AddNode add node to group
func AddNode(g Group, nd Node) {

}

//AddChildren add children to group
func AddChildren(g Group, ch []string) {

}

//AddVars add variables to group
func AddVars(g Group, vrs []string) {

}
