# simple-golang-grpc

#### 介绍

简单的 golang grpc 封装。支持 和 java 项目grpc 交互。支持 grpc 和 http 可以基于此二次开发、定制。
本项目是一个 golang项目 rpc服务提供者 脚手架。

#### 使用说明

1. 默认使用 dev 环境，本地运行需要创建 config.dev.yaml 配置文件。部署需要指定环境变量API_ENV:test/prod
2. go mod tidy
3. go build
4. ./main
