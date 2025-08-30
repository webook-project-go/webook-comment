//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/webook-project-go/webook-comment/grpc"
	"github.com/webook-project-go/webook-comment/ioc"
	"github.com/webook-project-go/webook-comment/repository"
	"github.com/webook-project-go/webook-comment/repository/cache"
	"github.com/webook-project-go/webook-comment/repository/dao"
	"github.com/webook-project-go/webook-comment/service"
)

var thirdPartyProvider = wire.NewSet(
	ioc.InitDatabase,
	ioc.InitRedis,
	ioc.InitEtcd,
	ioc.InitGrpcServer,
)

var commentServiceProvider = wire.NewSet(
	dao.NewDao,
	cache.NewRedisCache,
	repository.NewRepository,
	service.NewService,
)

func NewApp() *App {
	wire.Build(
		wire.Struct(new(App), "*"),
		thirdPartyProvider,
		commentServiceProvider,
		grpc.NewService,
	)
	return nil
}
