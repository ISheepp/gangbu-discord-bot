package server

import (
	"context"
	"encoding/json"
	"gangbu/pkg/models"
	"gangbu/pkg/queue"
	"gangbu/pkg/util"
	"gangbu/proto/game"
	"github.com/segmentio/kafka-go"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type gameServer struct {
	game.UnimplementedGameRequestServer // 如果一个结构体里只有类型没有指定字段名称，则这是一个嵌入类型
	ghUsecase                           models.GameHistoryUsecase
	pUsecase                            models.PlayerUsecase
}

func (g *gameServer) SendCallbackMessage(ctx context.Context, message *game.CallbackMessage) (*emptypb.Empty, error) {
	// i := reflect.ValueOf(message.Type).Interface()
	writer, err := queue.NewKafkaWriter("callback_msg")
	if err != nil {
		util.Logger.Error("创建kafka writer失败!", err)
		return nil, err
	}
	marshal, err := json.Marshal(message)
	if err != nil {
		util.Logger.Error("序列化失败!", err)
		return nil, err
	}
	util.Logger.Info("序列化完成，结果：", string(marshal))
	err = writer.WriteMessages(context.Background(), kafka.Message{Value: marshal})
	if err != nil {
		util.Logger.Error("写入kafka失败!", err)
		return nil, err
	}
	util.Logger.Info("写入kafka成功!")
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g *gameServer) CreateGame(ctx context.Context, bo *game.GameCreateBo) (*wrapperspb.StringValue, error) {
	ghu := g.ghUsecase
	resp, err := ghu.CreateGame(models.GameHistoryBo{
		PlayerDiscordUserId: bo.PlayerDiscordUserId,
		Choice:              uint8(bo.Choice),
		BetValue:            bo.BetValue,
		GuildID:             bo.GuildId,
		ChannelID:           bo.ChannelId,
	})
	if err != nil {
		util.Logger.Error("创建游戏失败!", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	util.Logger.Info("创建游戏成功!")
	return &wrapperspb.StringValue{Value: resp.Hash().Hex()}, nil
}

func (g *gameServer) GetLastFiveGameHistoryByDiscordId(ctx context.Context, discordId *wrapperspb.StringValue) (*game.GameHistoryDtoSlice, error) {
	ghu := g.ghUsecase
	gameHistory, err := ghu.GetLastFiveGameHistoryByDiscordId(discordId.GetValue())
	if err != nil {
		util.Logger.Error("获取游戏记录失败!", err)
		return nil, status.Errorf(codes.Internal, "获取游戏记录失败!")
	}
	var resultHistoryDto []*game.GameHistoryDto
	for _, history := range gameHistory {
		dto := &game.GameHistoryDto{
			GameResult: history.GameResult,
			BetValue:   history.BetValue,
			FinishTime: timestamppb.New(*history.FinishTime),
		}
		resultHistoryDto = append(resultHistoryDto, dto)
	}
	result := &game.GameHistoryDtoSlice{Histories: resultHistoryDto}
	return result, nil
}

func NewGameServer(ghu models.GameHistoryUsecase, pu models.PlayerUsecase) game.GameRequestServer {
	return &gameServer{
		ghUsecase: ghu,
		pUsecase:  pu,
	}
}
