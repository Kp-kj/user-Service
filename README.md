# user-service 
[![CodeQL](https://github.com/Kp-kj/user-Service/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/Kp-kj/user-Service/actions/workflows/github-code-scanning/codeql)[![Build and Push Docker Image](https://github.com/Kp-kj/user-Service/actions/workflows/go.yml/badge.svg)](https://github.com/Kp-kj/user-Service/actions/workflows/go.yml)[![Sync Develop to Release](https://github.com/Kp-kj/user-Service/actions/workflows/daily-to-release-sync.yml/badge.svg)](https://github.com/Kp-kj/user-Service/actions/workflows/daily-to-release-sync.yml)

## 项目结构
```
├── README.md
├── user.proto  // pb文件
├── user.go // 启动文件
├── userClient // go-zero生成的客户端代码
│   └── user.go
├── user        // go-zero生成的服务端pb
│   ├── user.pb.go
│   └── user_grpc.pb.go 
├── logs // 日志目录
├── internal // 内部包或模块
│   ├── config
│   │   └── config.go // 配置文件
│   ├── logic    // 业务逻辑
│   ├── model    // 数据库操作
│   │   └── ddl  // 数据库表结构
│   ├── server
│   └── svc      
└── etc
    └── user.yaml // 配置文件
```

## 项目启动
```
- 进入项目目录 user 

- go run user.go -f etc/user.yaml

注意：如果启动失败，可能是因为没有安装依赖，可以使用 go mod tidy 安装依赖
项目运行在 8080 端口,  如果需要调试，可以在 u ser.go 中修改端口
如需访问请先启动 etcd

```


## 项目常用命令
```
- 生成 rpc 代码
goctl rpc new user   //user-Service 目录运行

- 由ddl文件生成model操作代码
goctl model mysql ddl --src user.sql --dir ./myssql
注意,需要在ddl 目录运行. 生成的代码在 ddl的mysql目录下。
生成后将需要的代码放到 internal/model 目录下(为的是避免混淆)。
生成后记得删除ddl下的 mysql目录

- 设置etcd环境变量
set ETCDCTL_API=3

- windows下启动 etcd
.\etcdctl get user.rpc --prefix




```
注： 带缓存创建mysql的命令是 
goctl model mysql ddl --src user.sql --dir ./cache -cache


启动该容器在服务器上 

拉取 最新docker 镜像
docker pull fibonacci77777/userservice

启动容器 映射端口
docker run -d -p 8081:8080 --name userservice fibonacci77777/userservice


附： 会用到的docker命令
docker ps -a
docker ps  // 查看正在运行的容器
docker ps -a   // 查看所有容器
docker stop 98f7a8061434 // 停止容器
docker rm 98f7a8061434  // 删除容器
docker rmi 98f7a8061434  // 删除镜像


docker-compose up -d  // 后台启动容器
docker-compose stop <服务名称>
docker-compose rm <服务名称>
