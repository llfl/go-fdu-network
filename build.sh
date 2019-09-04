#！/bin/bash

# GOOS：目标平台的操作系统(darwin、freebsd、linux、windows)
# GOARCH：目标平台的体系架构(386、amd64、arm)
BUILD_TARGETS=('windows_amd64' 'windows_386' 'linux_amd64' 'linux_386' 'darwin_amd64')

if [ "$1" = 'clean_all' ] || [ "$1" = 'ca' ];then
    rm -rf ./release
elif [ "$1" = 'zip' ] || [ "$1" = 'z' ];then
    ZIPFLAG='Z'
elif [ "$1" = 'clean_zip' ] || [ "$1" = 'cz' ];then
    ZIPFLAG='CZ'
else
    BUILD_TARGETS=("$*")
fi



for TARGET in ${BUILD_TARGETS[*]}
do
    TARGET_OS=${TARGET%_*}
    TARGET_ARCH=${TARGET#*_}
    if [ $TARGET_OS = 'windows' ];then
        FILENAME='go-network.exe'
    else
        FILENAME='go-network'
    fi
    mkdir -p ./release/$TARGET
    GOOS=$TARGET_OS GOARCH=$TARGET_ARCH go build  -o ./release/$TARGET/$FILENAME ./go-network.go
    cp ./config.json ./release/$TARGET
    if [ $TARGET_OS = 'linux' ];then
        cp ./go-network.service ./release/$TARGET
    fi

    if [ $ZIPFLAG = 'Z' ];then
        zip -r ./release/$TARGET'.zip' ./release/$TARGET
    elif [ $ZIPFLAG = 'CZ' ];then
        zip -r -m ./release/$TARGET'.zip' ./release/$TARGET
    fi
done
