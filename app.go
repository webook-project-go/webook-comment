package main

import (
	"github.com/webook-project-go/webook-comment/grpc"
	"github.com/webook-project-go/webook-pkgs/grpcx"
)

type App struct {
	Server  *grpcx.GrpcxServer
	Service *grpc.Service
}
