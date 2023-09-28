package server

import (
	"context"
	"encoding/json"
	"fmt"
	myGraphql "gangbu/graphql"
	"gangbu/pkg/models"
	"gangbu/pkg/queue"
	"gangbu/pkg/util"
	"gangbu/proto/game"
	"github.com/segmentio/kafka-go"
	"github.com/shurcooL/graphql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"os"
	"strconv"
	"time"
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

func (g *gameServer) GetLastFiveGameHistoryByDiscordIdFromTheGraph(ctx context.Context, discordId *wrapperspb.StringValue) (*game.GameHistoryDtoSlice, error) {
	player, err := g.pUsecase.GetByDiscordUserID(discordId.GetValue())
	if err != nil {
		util.Logger.Error("查询玩家信息失败!", err)
		return nil, status.Errorf(codes.Internal, "query gamer failed!")
	}
	client := graphql.NewClient(os.Getenv("GRAPH_ENDPOINT_DEV"), nil)
	queryer := myGraphql.NewQueryer(client)
	var query struct {
		BetResults []struct {
			Id             graphql.String
			Amount         graphql.String
			RequestId      graphql.String
			CallerAddress  graphql.String
			GameResult     graphql.Boolean
			BlockTimestamp graphql.String
		} `graphql:"betResults(first:5 where: {callerAddress: $addr} orderBy: blockTimestamp orderDirection: desc)"`
	}
	vars := map[string]any{
		"addr": player.WalletAddress,
	}
	err = queryer.Query(context.Background(), &query, vars)
	if err != nil {
		util.Logger.Error("从the graph获取最近5条游戏记录失败!", err)
		return nil, status.Errorf(codes.Internal, "get data from the graph failed!")
	}
	var resultHistoryDto []*game.GameHistoryDto

	for _, history := range query.BetResults {
		betValue, err := strconv.ParseInt(string(history.Amount), 10, 64)
		if err != nil {
			util.Logger.Error("转换int失败!", err)
			return nil, status.Error(codes.Internal, "parse int failed!")
		}
		// 将字符串转换为整数类型
		timestamp, err := strconv.ParseInt(string(history.BlockTimestamp), 10, 64)
		if err != nil {
			fmt.Println("解析时间戳失败:", err)
			return nil, status.Error(codes.Internal, "parse timestamp failed!")
		}
		// 使用 time.Unix 将时间戳转换为 time.Time
		timestampTime := time.Unix(timestamp, 0)
		dto := &game.GameHistoryDto{
			GameResult: bool(history.GameResult),
			BetValue:   betValue,
			FinishTime: timestamppb.New(timestampTime),
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
