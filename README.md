Anywhere 随启随用的静态文件服务器
==============================

Running static file server anywhere. 随时随地将你的当前目录变成一个静态文件服务器的根目录。

## use

`fileserv -p port`,default port is 8001

1. direct use
```
go run fileserv.go [-p 8080]
```
2. build and run
```
go build src/fileserv.go
./fileserve [-p 8080]
```

3. executable file
```
Windows 64 bit:     fileserv.exe
Mac/Linux 64 bit:  ./fileserv
```

## Visit

another in the same local network
(同一局域网访问) 
```
http://ip:8001        //default,ip just try follow the tips which server given.
http://ip:<your port> //self defined port
```
or phone visit(in a same LAN):
```
http://<your server ip>:<port>
```

# WIFI使用
## 文件服务器最好的作用是在同一局域网(电脑开wifi共享给手机，共用一个路由器情况)下,手机无线浏览并下载电脑上的文件，支持多层目录访问
1. 在相应下载目录运行fileserv
2. windows运行`ipconfig /all`,Mac/Linux运行`ifconfig -a` 获取当前主机IP地址
3. 手机访问http://主机IP:端口号(IP由服务器显示给出,默认为8001)
4. 开始文件浏览下下载

# 测试结果
- 我的主机开wifi共享给手机，使用fileserv下载多个任务基本可以达到2M/s
- 不会出现像python服务器那样的中文乱码问题
- 不会出现python自带服务器半天无法下载和只能单个任务下载，后续任务长时间无响应等问题

