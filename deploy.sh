#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o openthailand-api main.go
git add .
git commit -m "deploy"
git push live master