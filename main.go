package main

import (
	"context"
	v1 "github.com/webook-project-go/webook-apis/gen/go/apis/comment/v1"
	_ "github.com/webook-project-go/webook-comment/config"
	"github.com/webook-project-go/webook-comment/ioc"
)

func main() {
	app := NewApp()
	shut := ioc.InitOTEL()
	defer shut(context.Background())
	v1.RegisterCommentServiceServer(app.Server.Server, app.Service)

	err := app.Server.Serve()
	if err != nil {
		panic(err)
	}
}
