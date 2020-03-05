# invcomb
(Ansible) inventory combinator tool for merging several inventories to one.

# install 

> go get github.com/peterducai/invcomb

> cd $GOPATH/src/github.com/peterducai/invcomb
> ./rebuild.sh

# use

> invcomb --input="examples/inventory1.yml,examples/inventory2.yml" --output="xxx.yml"

# Scenario


INFRASTRUCTURE team creates 10 different/separate environments (and their inventories) for several teams in company. They call them lab1 to lab10.

* They want to assign lab1 and lab2 to Security team.
* They want to assign lab3 to lab10 to development team.
* They want Security team to maintain all log servers in lab1 to lab10.
* They want security testers to be able pentest all webservers in lab1 to lab10.


 Instead of giving them lab1..10 inventory, they will create combined inventories with meaningful names like **security_dev_inventory** or **pentesting_lab**. Combining these inventories by hand could be error prone and with huge inventories (100 of groups) could be almost impossible, not just regarding time but also regarding avoidance of errors/typos. Another things is that for example pentester team, should not have info about rest of infrastructure, therefor giving them all lab1 to lab10 inventories make it more confusing for them (too many IPs) and less secure for others (they know other hosts!).



 <!-- By default variables are merged/flattened to the specific host before a play is run. This keeps Ansible focused on the Host and Task, so groups don’t really survive outside of inventory and host matching. By default, Ansible overwrites variables including the ones defined for a group and/or host (see DEFAULT_HASH_BEHAVIOUR). The order/precedence is (from lowest to highest):

all group (because it is the ‘parent’ of all other groups)
parent group
child group
host -->