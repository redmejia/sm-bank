#!/bin/bash

# clean term
clear 

echo "set db env"
source .dbenv

echo "running server"
go run *.*go


