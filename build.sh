#！/bin/bash

# GOOS：目标平台的操作系统(darwin、freebsd、linux、windows)
# GOARCH：目标平台的体系架构(386、amd64、arm)
BUILD_TARGETS=('windows_amd64' 'windows_386' 'linux_amd64' 'linux_386' 'darwin_amd64' 'linux_armv6' 'linux_armv7')

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
    if [ $TARGET_ARCH = 'armv7' ];then
        GOARM=7
        TARGET_ARCH='arm'
    elif [ $TARGET_ARCH = 'armv6' ];then
        GOARM=6
        TARGET_ARCH='arm'
    fi
    mkdir -p ./release/$TARGET
    GOOS=$TARGET_OS GOARCH=$TARGET_ARCH go build  -o ./release/$TARGET/$FILENAME ./go-network.go
    cp ./config.json ./release/$TARGET
    if [ $TARGET_OS = 'linux' ];then
        cp ./go-network.service ./release/$TARGET
        cp ./installation.sh ./release/$TARGET
    fi

    if [ $ZIPFLAG = 'Z' ];then
        cd ./release
        zip -r ./$TARGET'.zip' ./$TARGET
        cd ..
    elif [ $ZIPFLAG = 'CZ' ];then
        cd ./release
        zip -r -m ./$TARGET'.zip' ./$TARGET
        cd ..
    fi
done
