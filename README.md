# dilu-protobuf

## 目录结构
```
.
├── README.md
├── sh # 脚本
│   └── go_gen.sh # 自动生成protobuf文件
|
├── proto # protobuf文件目录
│   ├── mq # 消息队列
│   │   ├── dilu_protobuf.proto
│   │   └── dilu_protobuf.pb.go
│   ├── rpc # rpc服务
│   │   ├── dilu_protobuf.proto
│   │   └── dilu_protobuf.pb.go
│   └── socket # socket服务
│       ├── dilu_protobuf.proto
│       └── dilu_protobuf.pb.go

