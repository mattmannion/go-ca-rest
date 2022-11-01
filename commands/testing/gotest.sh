#!/bin/bash

cd src/api 

go test ./... -coverpkg=./... -coverprofile=../../coverage/coverage.out &&
clear &&
go tool cover -func ../../coverage/coverage.out 

# The testing configuration only covers the concrete implementation
# of the functions, methods, and branches therein that our program
# needs to run. Anything interchangable or configuation related is
# silently tested through a successful compile and run of the server
# and testing suite. Anything that needs to be mocked, will be, in src/api.
# Use clients/pg/* for an example of an interchangable dependency.
