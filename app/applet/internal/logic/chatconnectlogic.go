package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/chat/pb/chat"

	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

// 全局连接管理
var (
	clients    = make(map[int64]*websocket.Conn)
	clientsMux sync.RWMutex
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatConnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
	userId int64
}

func NewChatConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request, userId int64) *ChatConnectLogic {
	return &ChatConnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
		userId: userId,
	}
}

func (l *ChatConnectLogic) ChatConnect() error {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(l.w, l.r, nil)
	if err != nil {
		return fmt.Errorf("websocket upgrade error: %v", err)
	}
	defer conn.Close()
	// 保存连接
	clientsMux.Lock()
	clients[l.userId] = conn
	clientsMux.Unlock()

	// 清理连接
	defer func() {
		clientsMux.Lock()
		delete(clients, l.userId)
		clientsMux.Unlock()
	}()
	// 处理连接
	for {
		// 读取消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Errorf("websocket read error: %v", err)
			}
			break
		}

		// 解析消息
		var req types.ChatRequest
		if err := json.Unmarshal(message, &req); err != nil {
			logx.Errorf("json unmarshal error: %v", err)
			continue
		}

		// 处理消息类型
		switch req.Type {
		case "message":
			// 构建消息
			msg := types.ChatMessage{
				Id:         time.Now().UnixNano(),
				SenderId:   l.userId,
				ReceiverId: req.ReceiverId,
				Content:    req.Content,
				CreateTime: time.Now().Unix(),
			}

			// 发送消息到接收者
			if err := l.sendToUser(req.ReceiverId, "message", msg); err != nil {
				logx.Errorf("send message error: %v", err)
			}

			// 保存消息到数据库
			if err := l.saveMessage(msg); err != nil {
				logx.Errorf("save message error: %v", err)
			}

		case "heartbeat":
			// 响应心跳
			if err := l.sendToUser(l.userId, "heartbeat", types.ChatMessage{}); err != nil {
				logx.Errorf("send heartbeat error: %v", err)
			}
		}
	}
	return nil
}

// 发送消息给指定用户
func (l *ChatConnectLogic) sendToUser(userId int64, msgType string, message types.ChatMessage) error {
	clientsMux.RLock()
	conn, ok := clients[userId]
	clientsMux.RUnlock()

	if !ok {
		return fmt.Errorf("user %d not online", userId)
	}

	resp := types.ChatResponse{
		Type:    msgType,
		Message: message,
	}

	return conn.WriteJSON(resp)
}

// 保存消息到数据库
func (l *ChatConnectLogic) saveMessage(msg types.ChatMessage) error {
	// 调用chat服务保存消息
	_, err := l.svcCtx.ChatRPC.SendMassage(l.ctx, &chat.SendMassageRequest{
		SenderId:   msg.SenderId,
		ReceiverId: msg.ReceiverId,
		Content:    msg.Content,
	})
	return err
}
