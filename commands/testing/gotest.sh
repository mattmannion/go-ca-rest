#!/bin/bash

cd cmd/src 
go test ./... -v -coverpkg=./... -covermode=count