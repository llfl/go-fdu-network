#！/bin/bash


#GOOS：目标平台的操作系统（darwin、freebsd、linux、windows） 
#GOARCH：目标平台的体系架构（386、amd64、arm） 
GOOS=windows GOARCH=amd64 go build ./go-network.go