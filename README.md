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
