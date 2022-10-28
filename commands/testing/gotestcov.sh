#!/bin/bash

cd cmd/src 
go test ./... -coverpkg=./... -coverprofile=../../coverage/coverage.out
go tool cover -func ../../coverage/coverage.out 