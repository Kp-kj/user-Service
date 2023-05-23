# user-service

set ETCDCTL_API=3

.\etcdctl get user.rpc --prefix

//  生成 rpc 代码
goctl rpc proto -csrc rpc/user/user.proto -dir .