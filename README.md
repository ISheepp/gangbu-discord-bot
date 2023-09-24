# gangbu-discord-bot

깐부: 类似 Squid 里的猜单双游戏

## discord commands

+ /play : 进行游戏
+ /show : 展示当前用户的余额，充值地址
+ /history: 查看游戏历史记录
+ /withdraw : 提款 「输入address，amount」

## 步骤

1. 先展示用户信息来生成地址，然后往里面充值
2. 使用 /play 来进行游戏「输入您的选择和筹码」
3. 等待 bot 返回游戏结果

## 设计


## 开发时遇到的问题

[如何由合约abi生成go文件](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings)

`abigen --abi abi.abi --pkg models --type EvenOddGame --out even_odd_game.go`

[参考 gorm 事务](https://blogs.halodoc.io/db-transactions-in-go/)

[kafka-go 连接问题](https://github.com/segmentio/kafka-go/issues/499)

## grpc相关

[protoc3语法](https://protobuf.dev/programming-guides/proto3/)

1. 在编写proto文件的时候，因为定义rpc必须要有参数和返回体，我们可以导入google事先定义好的一些包装类型
google/protobuf/wrappers.proto，空类型google/protobuf/empty.proto
2. 生成pb.go文件，其中game.pb.go文件包含了message的类型，game_grpc_pb.go包含了rpc方法的定义，服务端需要实现此方法，客户端使用其中的接口方法

## warning

discord slash command 只有3s的时间给后端响应