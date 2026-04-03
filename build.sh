#!/bin/bash

GOOS=windows GOARCH=amd64 go build 
tar -czf migo-Windows-amd64.tar.gz mi.exe

GOOS=linux GOARCH=amd64 go build
tar -czf migo-Linux-amd64.tar.gz mi
