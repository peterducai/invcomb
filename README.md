# invcomb
(Ansible) inventory combinator tool for merging several inventories to one.

# install 

> go get github.com/peterducai/invcomb


# use

> invcomb --input="examples/inventory1.yml,examples/inventory2.yml" --output="xxx.yml"

# Scenario


INFRA team creates 10 different labs with separate networks/IPs.

They will assign lab1 and lab2 for team T1. Instead of giving them lab1/2 inventory, they will create combined T1_inventory.