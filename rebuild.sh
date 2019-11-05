#!/bin/bash

go build
go install
clear 
invcomb --input="examples/inventory1.yml,examples/inventory2.yml,examples/inventory3.yml"