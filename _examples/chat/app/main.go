package main

import (
	"github.com/cherry-game/cherry"
	"github.com/cherry-game/cherry/_examples/chat"
	"github.com/cherry-game/cherry/component/gin"
	"github.com/cherry-game/cherry/net/connector"
	"github.com/cherry-game/cherry/net/handler"
	"github.com/cherry-game/cherry/net/serializer"
	"github.com/cherry-game/cherry/net/session"
)

func main() {
	app := cherry.NewApp("../profile/", "local", "gate-1")
	app.SetSerializer(cherrySerializer.NewJSON())

	httpServer := cherryGin.New("127.0.0.1:80", cherryGin.RecoveryWithZap(true))
	httpServer.StaticFS("/", "../web/")

	wsComponent := cherryConnector.NewWSComponent("127.0.0.1:34590")

	app.Startup(
		cherrySession.NewComponent(),
		handlerComponent(),
		httpServer,
		wsComponent,
	)
}

func handlerComponent() *cherryHandler.Component {
	component := cherryHandler.NewComponent()

	group1 := cherryHandler.NewGroup(10, 256)
	group1.AddHandlers(&chat.UserHandler{})
	component.Register(group1)

	group2 := cherryHandler.NewGroup(10, 256)
	group2.AddHandlers(&chat.RoomHandler{})
	component.Register(group2)

	return component
}
