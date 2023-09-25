package main

import (
	"context"
	"fmt"
	"gangbu/game/listener"
	_gameRepo "gangbu/game/repository/mysql"
	_gameServer "gangbu/game/server"
	_gameUsecase "gangbu/game/usecase"
	"gangbu/pkg/db"
	"gangbu/pkg/util"
	_playerRepo "gangbu/player/repository/mysql"
	_playerServer "gangbu/player/server"
	_playerUsecase "gangbu/player/usecase"
	"gangbu/proto/game"
	"gangbu/proto/player"
	"gangbu/proto/withdraw"
	_withdrawServer "gangbu/withdraw/server"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
)

// RunGrpcServer 启动Grpc服务
func RunGrpcServer(wg *sync.WaitGroup) {
	util.Logger.Info("start grpc server")
	dbClient := db.NewDBClient(context.Background())
	playerRepository := _playerRepo.NewPlayerRepository(dbClient)
	gameHistoryRepository := _gameRepo.NewGameHistoryRepository()
	gameUsecase := _gameUsecase.NewGameUsecase(gameHistoryRepository, playerRepository, dbClient)
	playerUsecase := _playerUsecase.NewPlayerUsecase(playerRepository)
	gameServer := _gameServer.NewGameServer(gameUsecase, playerUsecase)
	playerServer := _playerServer.NewPlayerServer(playerUsecase)
	withdrawServer := _withdrawServer.NewWithdrawServer(playerUsecase)
	// 监听器
	chainListener := listener.NewChainListener(gameUsecase, playerUsecase)
	kafkaListener := listener.NewKafkaListener(playerUsecase)
	chainListener.StartListen()
	kafkaListener.StartListen()
	startGrpcServer(wg, gameServer, playerServer, withdrawServer)
}

// StartGrpcServer 注册Grpc服务
func startGrpcServer(wg *sync.WaitGroup, grs game.GameRequestServer, prs player.PlayerRequestServer, wrs withdraw.WithdrawRequestServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", os.Getenv("PORT")))
	if err != nil {
		util.Logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 注册服务
	game.RegisterGameRequestServer(s, grs)
	player.RegisterPlayerRequestServer(s, prs)
	withdraw.RegisterWithdrawRequestServer(s, wrs)
	// 启动RPC并监听
	util.Logger.Printf("server listening at %v", lis.Addr())
	wg.Done()
	if err = s.Serve(lis); err != nil {
		util.Logger.Fatalf("failed to serve: %v", err)
	}
}
