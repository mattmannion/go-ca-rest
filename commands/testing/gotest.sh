#!/bin/bash

cd cmd/api 

go test ./... -coverpkg=./... -coverprofile=../../coverage/coverage.out &&
go tool cover -func ../../coverage/coverage.out 