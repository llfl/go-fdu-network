#！/bin/bash

BUILD_TARGETS=('windows_amd64' 'windows_386' 'linux_amd64' 'darwin_amd64')

for TARGET in BUILD_TARGETS 
do
    TARGET_OS 
done

#GOOS：目标平台的操作系统（darwin、freebsd、linux、windows） 
#GOARCH：目标平台的体系架构（386、amd64、arm） 
GOOS=windows GOARCH=amd64 go build  -o ./release/Windows_amd64/go-network.exe ./go-network.go
cp ./config.json ./release/Windows_amd64/
zip -r ./release/Windows_amd64

GOOS=windows GOARCH=386 go build -o ./release/Windows_x86/go-network.exe ./go-network.go 
cp ./config.json ./release/Windows_x86/
zip -r ./release/Windows_x86

GOOS=linux GOARCH=amd64 go build -o ./release/Linux_amd64/go-network ./go-network.go
cp ./config.json ./release/Linux_amd64/
cp ./go-network.service ./release/Linux_amd64/
zip -r ./release/Linux_amd64

GOOS=darwin GOARCH=amd64 go build -o ./release/Darwin_amd64/go-network ./go-network.go
cp ./config.json ./release/Darwin_amd64/
zip -r ./release/Darwin_amd64