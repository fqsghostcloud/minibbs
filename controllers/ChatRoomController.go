package controllers

import (
	"container/list"
	"encoding/json"
	"fmt"
	"minibbs/filters"
	"minibbs/models"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type ChatRoomController struct {
	beego.Controller
}

// Get method handles GET requests for WebSocketController.
func (c *ChatRoomController) ChatRoomPage() {

	uname := c.GetString("uname")
	tid := c.GetString("tid")
	if len(uname) == 0 || len(tid) == 0 {
		c.Redirect("/", 302)
		return
	}

	//check usename is current user
	_, currUser := filters.IsLogin(c.Ctx)
	if currUser.Username != uname {
		c.Redirect("/", 302)
		return
	}

	topicId, err := strconv.Atoi(tid)
	if err != nil {
		beego.Error("get topic id for chatroon error[%s]", err.Error())
		c.Ctx.WriteString("发生错误，请联系管理员")
	}

	topic := models.TopicManager.FindTopicById(topicId)

	c.Data["UserInfo"] = currUser
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "聊天室"
	c.Data["UserName"] = uname
	c.Data["TopicName"] = topic.Title
	c.Data["TopicId"] = topicId
	c.Layout = "layout/layout.tpl"
	c.TplName = "chatroomcontroller/chatRoomPage.tpl"
	
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ChatRoomScript"] = "chatroomcontroller/ChatRoomScript.tpl"
}

// Join method handles WebSocket requests for WebSocketController.
func (c *ChatRoomController) Chat() {
	uname := c.GetString("uname")
	tid := c.GetString("tid")
	if len(uname) == 0 || len(tid) == 0 {
		c.Redirect("/", 302)
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
			return
		} else {
			beego.Error("Cannot setup WebSocket connection:", err)
			return
		}
	}

	if _, isExist := chatroomMap[tid]; !isExist {
		chatroomch := &ChatRoomCh{
			// Channel for new join users.
			comeinChatterCh: make(chan Chatter, 10),
			// Channel for exit users.
			exitChatterCh: make(chan string, 10),
			// Send events here to commonInfoCh them.
			commonInfoCh: make(chan models.Event, 10),

			chatterLists: list.New(),
		}

		chatroomMap[tid] = chatroomch
		go chatroom(chatroomMap[tid])
	}

	chatter := Chatter{Name: uname, TopicId: tid, Conn: ws}

	// Join chat room.
	Join(chatroomMap[tid], chatter)
	defer Leave(uname, tid)

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("\nwebsocket message receive loop erro[%s]\n", err.Error())
			return
		}

		chatroomMap[tid].commonInfoCh <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}

	c.TplName = "blank.tpl"

}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event, pChatroomch *ChatRoomCh) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for chatterItem := pChatroomch.chatterLists.Front(); chatterItem != nil; chatterItem = chatterItem.Next() {
		// Immediately send event to WebSocket users.
		ws := chatterItem.Value.(Chatter).Conn //断言
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				pChatroomch.exitChatterCh <- chatterItem.Value.(Chatter).Name
			}
		}
	}
}

func newEvent(ep models.EventType, user, msg string) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg}
}

func Join(chatroomch *ChatRoomCh, chatter Chatter) {
	chatroomch.comeinChatterCh <- chatter
}

func Leave(user string, topicId string) {
	chatroomMap[topicId].exitChatterCh <- user
}

type Chatter struct {
	Name    string
	TopicId string
	Conn    *websocket.Conn // Only for WebSocket users; otherwise nil.
}

type ChatRoomCh struct {
	comeinChatterCh chan Chatter
	exitChatterCh   chan string
	commonInfoCh    chan models.Event
	chatterLists    *list.List
}

var chatroomMap = make(map[string]*ChatRoomCh)

// This function handles all incoming chan messages.
func chatroom(v *ChatRoomCh) {
	for {
		select {
		case chatter := <-v.comeinChatterCh:
			if !isUserExist(v.chatterLists, chatter.Name) {
				v.chatterLists.PushBack(chatter) // Add user to the end of list.
				// Publish a JOIN event.
				v.commonInfoCh <- newEvent(models.EVENT_JOIN, chatter.Name, "")
				beego.Info("New user:", chatter.Name, ";WebSocket:", chatter.Conn != nil)
			} else {
				beego.Info("Old user:", chatter.Name, ";WebSocket:", chatter.Conn != nil)
			}
		case event := <-v.commonInfoCh:

			broadcastWebSocket(event, v)
			// models.AddEvent(event) 从events list 获取消息历史记录

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case unsub := <-v.exitChatterCh:
			for sub := v.chatterLists.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Chatter).Name == unsub {
					v.chatterLists.Remove(sub)
					// Clone connection.
					ws := sub.Value.(Chatter).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					v.commonInfoCh <- newEvent(models.EVENT_LEAVE, unsub, "") // Publish a LEAVE event.
					break
				}
			}
		}

	}
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Chatter).Name == user {
			return true
		}
	}
	return false
}
