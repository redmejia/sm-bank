#!/bin/bash

# clean term
clear 

echo "set db env"
source .dbenv
echo "...done"

echo "Run server"
go run *.*go


