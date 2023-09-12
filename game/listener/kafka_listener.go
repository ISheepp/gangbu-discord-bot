package listener

import (
	"context"
	"gangbu/pkg/queue"
	"gangbu/pkg/util"
)

type kafkaListener struct {
}

func NewKafkaListener() Listener {
	return &kafkaListener{}
}

func (k *kafkaListener) StartListen() {
	go ListenGameDoneHistory()
}

func ListenGameDoneHistory() {
	reader, err := queue.NewKafkaReader()
	if err != nil {
		util.Logger.Error("创建kafka reader失败", err)
	}
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			util.Logger.Error("kafka reader读取消息失败", err)
			continue
		}
		util.Logger.Info("kafka reader读取消息成功\n", string(msg.Value))
		// todo 推送到discord

	}
}
