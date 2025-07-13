package main

import (
	"LingXi/app/chat/cmd/api/internal/config"
	"LingXi/app/chat/cmd/api/internal/handler"
	"LingXi/app/chat/cmd/api/internal/svc"
	"LingXi/app/chat/cmd/rpc/chatcenter"
	chat1 "LingXi/common/chat"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"log"
	"net/http"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	go Run(ctx)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func Run(serverCtx *svc.ServiceContext) {
	r := &http.Request{}
	for {
		select {
		case client := <-serverCtx.Hub.Register:
			serverCtx.Hub.Clients[client] = true
		case client := <-serverCtx.Hub.Unregister:
			if _, ok := serverCtx.Hub.Clients[client]; ok {
				delete(serverCtx.Hub.Clients, client)
				close(client.Send)
			}
		case message := <-serverCtx.Hub.Broadcast:
			for client := range serverCtx.Hub.Clients {
				var msg chat1.Message
				err := json.Unmarshal(message, &msg)
				if err != nil {
					log.Println("Failed to unmarshal message:", err)
					break
				}
				//fmt.Println(client.UserId)
				//fmt.Println("----------------")
				//fmt.Println(msg)
				if msg.ReceiverId == client.UserId {
					select {
					case client.Send <- message:
						serverCtx.ChatCenterRpc.SaveMessageRequest(r.Context(), &chatcenter.SaveMessageRequest{
							SenderId:   msg.SenderId,
							ReceiverId: msg.ReceiverId,
							Content:    msg.Content,
							Type:       msg.Type,
						})
					default:
						close(client.Send)
						delete(serverCtx.Hub.Clients, client)
					}
				}
			}
		}
	}

}
