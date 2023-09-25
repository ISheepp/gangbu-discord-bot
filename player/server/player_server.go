package server

import (
	"context"
	"gangbu/pkg/models"
	"gangbu/pkg/util"
	"gangbu/proto/player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type playerServer struct {
	player.UnimplementedPlayerRequestServer
	pUsecase models.PlayerUsecase
}

func (p *playerServer) ShowPlayerInfo(ctx context.Context, discordId *wrapperspb.StringValue) (*player.PlayerInfoVo, error) {
	pu := p.pUsecase
	playPerson, err := pu.GetByDiscordUserIDOrCreate(discordId.GetValue())
	if err != nil {
		util.Logger.Error("查询玩家信息失败！", err)
		return nil, status.Errorf(codes.Internal, "查询玩家信息失败!, id%s", discordId.GetValue())
	}
	result := &player.PlayerInfoVo{
		Id:            uint64(playPerson.ID),
		DiscordUserId: playPerson.DiscordUserId,
		WalletAddress: playPerson.WalletAddress,
		WalletValue:   playPerson.WalletValue,
		CreateAt:      timestamppb.New(playPerson.CreateAt),
	}
	return result, nil
}

func NewPlayerServer(pu models.PlayerUsecase) player.PlayerRequestServer {
	return &playerServer{
		pUsecase: pu,
	}
}
