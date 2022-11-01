#!/bin/bash

cd src/api 

go test ./... -coverprofile=../../coverage/coverage.out &&
clear &&
go tool cover -func ../../coverage/coverage.out 