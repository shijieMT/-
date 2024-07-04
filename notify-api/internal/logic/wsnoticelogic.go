package logic

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"notify/common"
	"time"

	"notify/notify-api/internal/svc"
	"notify/notify-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WSNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWSNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WSNoticeLogic {
	return &WSNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 发送消息到客户端
func sendMessage(conn *websocket.Conn, messageType int, message []byte) error {
	conn.SetWriteDeadline(time.Now().Add(writeWait))
	return conn.WriteMessage(messageType, message)
}

func (l *WSNoticeLogic) WSNotice(req *types.WSNoticeRequest, w http.ResponseWriter, r *http.Request) (resp *types.WSNoticeResponse, err error) {
	// todo: add your logic here and delete this line

	// 验证
	userid, err := common.Auth(req.Token)
	if err != nil {
		return nil, err
	}
	_ = userid

	// 升级 HTTP 连接至 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// 设置读写超时
	conn.SetReadDeadline(time.Now().Add(pongWait)) // 设置初始读超时为60秒
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait)) // 每次收到pong消息，更新读超时
		return nil
	})

	ticker := time.NewTicker(pingPeriod) // 每隔54秒发送一次ping消息
	defer ticker.Stop()

	go func() {
		for {
			select {
			case t := <-ticker.C:
				if err := sendMessage(conn, websocket.PingMessage, nil); err != nil {
					log.Println("PingMessage error:", err)
					return
				}
				log.Println("Ping sent at:", t)
			}
		}
	}()

	// todo 读取和处理客户端消息的部分
	// todo 使用数据库通知机制或触发器+消息队列
	for {
		// 读取客户端消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error:", err)
			break
		}
		log.Printf("Received message: %s", message)

		// 回应客户端消息
		response := fmt.Sprintf("Echo: %s", message)
		if err := sendMessage(conn, websocket.TextMessage, []byte(response)); err != nil {
			log.Println("WriteMessage error:", err)
			break
		}
	}

	return
}

// 读写超时时间
const (
	writeWait  = 10 * time.Second    // 写操作超时时间为10秒
	pongWait   = 60 * time.Second    // 等待客户端pong消息的时间为60秒
	pingPeriod = (pongWait * 9) / 10 // 发送ping消息的周期为54秒（即60秒的90%）
)

// go get github.com/gorilla/websocket
// 升级器配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
