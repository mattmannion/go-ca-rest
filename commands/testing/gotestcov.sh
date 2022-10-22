#!/bin/bash

cd cmd/src 
go test ./... -v -coverpkg=./... -coverprofile=../../coverage/coverage.out
go tool cover -func ../../coverage/coverage.out 