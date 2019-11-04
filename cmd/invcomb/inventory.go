package invcomb

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
var Inv = new(Inventory)
var emptyVar = new(Variable)
var emtpyGroup = new(Group)
var emptyNode = new(Node)

//InitInventory initialize inventory
func InitInventory() {
	emtpyGroup.Name = "x"
	emtpyGroup.Nodes = append(emtpyGroup.Nodes, *emptyNode)
	emtpyGroup.Children = append(emtpyGroup.Nodes, [])

	Inv.Author = "my author"
	Inv.Date = "today"
	Inv.Groups = append(Inv.Groups, *emtpyGroup)
	Inv.Nodes = append(Inv.Groups, *emptyNode)
	Inv.Variables = append(Inv.Groups, *emptyVar)
}

//AddGroup add group to inventory
func AddGroup(g string) {
	// var gr = new(Group)
	// gr.Name = g
	// gr.Nodes = append(gr.Nodes, nil)
	// gr.Children = nil
	// gr.Variables = nil
	// Inv = append(Inv, *gr)
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
