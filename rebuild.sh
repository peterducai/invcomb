#!/bin/bash

go build invcomb.go
go install

invcomb --inputfile="examples/inventory1.yml,examples/inventory2.yml,examples/inventory3.yml"  --outputfile="xxx.yml"
