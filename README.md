# go-fdu-network

> 该版本的go-fdu-network通过读取配置文件获取登录的用户名和密码 二进制包发布于release

## 编译

该版本优化了编译，可同时交叉编译多个平台的二进制文件。编译步骤如下

1. 在本目录下运行build.sh即可默认编译Windows、Linux、macOS平台的二进制文件并生成于./release 文件夹下。

2. build.sh 支持参数编译，如运行"./build.sh clean_all"或"./build.sh ca"可以清理./release。

``` notice
!!! 注意 !!!:

"build.sh clean_all" 包含 "rm -rf" 请注意不要使用管理员或任何root权限运行。
```

3. 直接输入相应的交叉编译参数即可编译相应的平台。交叉编译参数示例如下

```notice
"目标操作系统"+"下划线"+"目标构架",如：

若想为运行linux操作系统的arm64交叉编译 编译参数为"linux_arm64"

运行 "./build.sh linux_arm64"即可编译相应的二进制文件
```

## 使用

在config.json中写入用户名密码并放于go-network同目录下 运行./go-network 即可

### Windows

在Windows下用命令行直接运行go-network.exe即可

## Linux 开机自启

本版本配置了linux的systemd服务脚本go-network.service，按步骤安装即可。

> go-networkd.service 默认执行/usr/local/bin下的go-network 配置文件默认于/usr/local/src 要自定义配置文件以及执行地址请修改 go-networkd.service

1. 将编译出的二进制文件go-network 拷贝至/usr/local/bin目录下

2. 将config.json 拷贝至/usr/local/src目录下

3. 将go-networkd.service 拷贝至/etc/systemd/system目录下，并执行 systemctl enable go-network 即可设置go-network开机自启。
