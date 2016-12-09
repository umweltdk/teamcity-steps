#!/bin/bash
GOOS=linux   GOARCH=amd64 go-wrapper download
GOOS=windows GOARCH=386   go-wrapper download
GOOS=linux   GOARCH=amd64 go-wrapper install
GOOS=windows GOARCH=386   go-wrapper install
mkdir -p dist/linux_amd64/
cp /go/bin/teamcity-steps dist/linux_amd64/
cp -R /go/bin/windows_386/* dist/windows_386/