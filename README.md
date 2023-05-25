# user-service

set ETCDCTL_API=3

.\etcdctl get user.rpc --prefix

//  生成 rpc 代码
goctl rpc new user   //user-Service 目录运行

goctl model mysql ddl --src user.sql --dir ./cache -cache