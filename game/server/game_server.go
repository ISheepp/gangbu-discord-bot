package server

import (
	"context"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"gangbu/proto"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type gameServer struct {
	proto.UnimplementedGameRequestServer // 如果一个结构体里只有类型没有指定字段名称，则这是一个嵌入类型
	ghUsecase                            models.GameHistoryUsecase
	pUsecase                             models.PlayerUsecase
}

func NewGameServer(ghu models.GameHistoryUsecase, pu models.PlayerUsecase) *gameServer {
	return &gameServer{
		ghUsecase: ghu,
		pUsecase:  pu,
	}
}

func (g *gameServer) mustEmbedUnimplementedGameRequestServer() {

}

func (g *gameServer) CreateGame(ctx context.Context, bo *proto.GameCreateBo) (*emptypb.Empty, error) {

	panic("implement me")
}

func (g *gameServer) GetLastFiveGameHistoryByDiscordId(ctx context.Context, discordId *wrapperspb.StringValue) (*proto.GameHistoryDtoSlice, error) {
	ghu := g.ghUsecase
	gameHistory, err := ghu.GetGameHistoryByDiscordId(discordId.GetValue())
	if err != nil {
		util.Logger.Error("获取游戏记录失败!", err)
		return nil, status.Errorf(codes.Internal, "获取游戏记录失败!")
	}
	var resultHistoryDto []*proto.GameHistoryDto
	for _, history := range gameHistory {
		dto := &proto.GameHistoryDto{
			GameResult: history.GameResult,
			BetValue:   history.BetValue,
			// todo history.FinishTime如果为空则不能使用*取值
			FinishTime: timestamppb.New(*history.FinishTime),
		}
		resultHistoryDto = append(resultHistoryDto, dto)
	}
	result := &proto.GameHistoryDtoSlice{Histories: resultHistoryDto}
	return result, status.New(codes.OK, "获取游戏记录成功！").Err()
}
