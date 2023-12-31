package websockethandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mtsaudio/internal/tokens"
	"mtsaudio/internal/utils/httputils"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gorilla/websocket"
)

type WebsocketUsecase interface {
}

type WebsocketHandler struct {
	wu WebsocketUsecase
}

func New(wu WebsocketUsecase) WebsocketHandler {
	ChatRooms = make(map[string]Room)
	FileRooms = make(map[string]Room)
	TrackRooms = make(map[string]Room)

	return WebsocketHandler{wu}
}

type Client struct {
	username string
	conn     *websocket.Conn
}

type Room struct {
	OwnerId uint
	Clients map[Client]struct{}
	Mu      *sync.RWMutex
}

var ChatRooms map[string]Room
var FileRooms map[string]Room
var TrackRooms map[string]Room

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024,
	WriteBufferSize: 1024 * 1024,
	//Solving cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type chatMessage struct {
	Username string `json:"username"`
	Msg      string `json:"msg"`
}

type accessToken struct {
	AccessToken string `json:"access_token"`
}

func (wh WebsocketHandler) ChatConnect(roomId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer conn.Close()

		var token accessToken
		if err := conn.ReadJSON(&token); err != nil {
			log.Printf("failed to parse JSON: %s", err.Error())
			conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "access token required"))
			return
		}

		tokenData, err := tokens.ParseToken(token.AccessToken)
		if err != nil {
			conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "invalid access token"))
			log.Printf("failed to parse token")
			return
		}

		username := tokenData.Username

		client := Client{
			conn:     conn,
			username: username,
		}

		room := ChatRooms[roomId]
		defer func() {
			room.Mu.Lock()
			delete(room.Clients, client)
			room.Mu.Unlock()
			log.Printf("user: %s disconnected from chat room: %s", client.username, roomId)
		}()

		room.Mu.Lock()
		room.Clients[client] = struct{}{}
		room.Mu.Unlock()

		room.Mu.RLock()
		clients := make([]string, 0, len(room.Clients))
		for cl := range room.Clients {
			clients = append(clients, cl.username)
		}
		room.Mu.RUnlock()

		if err := conn.WriteJSON(clients); err != nil {
			conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "failed to send users list"))
			log.Println("failed to send user amount")
			return
		}

		log.Printf("user %s connected to chat room: %s", username, roomId)

		for {
			var msg chatMessage

			if err := conn.ReadJSON(&msg); err != nil {
				log.Printf("chat websocket error: %s", err.Error())
				continue
			}
			msg.Username = client.username

			log.Printf("chat message recived: roomId: %s, user: %s, msg: %s", roomId, msg.Username, msg.Msg)
			msgData, err := json.Marshal(&msg)
			if err != nil {
				log.Printf("chat websocket error: %s", err.Error())
				continue
			}
			go room.broadcast(msgData)
		}
	}
}

type fileMessage struct {
	Filename string `json:"file_name"`
}

type audioMessage struct {
	AudioName string `json:"audio_name"`
	Chunk     []byte `json:"chunk"`
	Done      bool   `json:"done"`
}

func (wh WebsocketHandler) FileConnect(roomId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer conn.Close()

		var token accessToken
		if err := conn.ReadJSON(&token); err != nil {
			log.Printf("failed to parse JSON: %s", err.Error())
			conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "access token required"))
			return
		}

		tokenData, err := tokens.ParseToken(token.AccessToken)
		if err != nil {
			conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "invalid access token"))
			log.Printf("failed to parse token")
			return
		}

		client := Client{
			conn:     conn,
			username: tokenData.Username,
		}

		room := FileRooms[roomId]

		defer func() {
			room.Mu.Lock()
			delete(room.Clients, client)
			room.Mu.Unlock()
			log.Printf("user: %s disconnected from file room: %s", client.username, roomId)
		}()

		room.Mu.Lock()

		room.Clients[client] = struct{}{}
		room.Mu.Unlock()
		log.Printf("user: %s connected to file room: %s", client.username, roomId)
		//listen messages
		for {
			var msg fileMessage
			if err := conn.ReadJSON(&msg); err != nil {
				continue
			}

			if tokenData.Id != room.OwnerId {
				log.Printf("no right to start sending file: userId = %d, ownerId = %d", tokenData.Id, room.OwnerId)
				continue
			}

			filePath := fmt.Sprintf("./static/%s", msg.Filename)
			absFilePath, err := filepath.Abs(filePath)
			if err != nil {
				log.Printf("failed to resolve abs path to file: %s", absFilePath)
				continue
			}

			fs, err := os.Open(absFilePath)

			if err != nil {
				log.Printf("failed to open file: %s", absFilePath)
				conn.WriteJSON(httputils.ResponseError{
					Error: fmt.Sprintf("failed to open file: %s", absFilePath),
				})
				continue
			}
			defer fs.Close()
			log.Printf("sending file: roomId = %s, filename = %s", roomId, msg.Filename)
			//sending file
			for {
				buffer := make([]byte, 1024)
				n, err := fs.Read(buffer)

				var done bool
				if n != 1024 {
					done = true
				}

				if err != nil {
					fs.Close()
					if err == io.EOF {
						log.Printf("finished to send file: %s", msg.Filename)
						break
					}
					log.Print("file websocket error: ", err.Error())
					break
				}

				msg := audioMessage{
					AudioName: msg.Filename,
					Chunk:     buffer[:n],
					Done:      done,
				}

				msgData, err := json.Marshal(msg)
				if err != nil {
					log.Printf("file websocket error: %s", err.Error())
					continue
				}

				room.broadcast(msgData)
			}
		}
	}
}

func (wh WebsocketHandler) TrackConnect(roomId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer conn.Close()

		var token accessToken
		if err := conn.ReadJSON(&token); err != nil {
			log.Printf("failed to parse JSON: %s", err.Error())
			conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "access token required"))
			return
		}

		tokenData, err := tokens.ParseToken(token.AccessToken)
		if err != nil {
			conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "invalid access token"))
			log.Printf("failed to parse token")
			return
		}

		client := Client{
			conn:     conn,
			username: tokenData.Username,
		}

		room := TrackRooms[roomId]

		defer func() {
			room.Mu.Lock()
			delete(room.Clients, client)
			room.Mu.Unlock()
			log.Printf("user: %s disconnected from track websocket: %s", client.username, roomId)
		}()

		room.Mu.Lock()

		room.Clients[client] = struct{}{}
		room.Mu.Unlock()
		log.Printf("user: %s connected to track websocket: %s", client.username, roomId)

		type musicPlay struct {
			PlayMusic bool `json:"play_music"`
		}

		for {
			var msg musicPlay
			if err := conn.ReadJSON(&msg); err != nil {
				continue
			}

			if tokenData.Id != room.OwnerId {
				log.Printf("room's owner access only")
				continue
			}

			msgData, err := json.Marshal(msg)
			if err != nil {
				log.Printf("chat websocket error: %s", err.Error())
				continue
			}

			go room.broadcast(msgData)
		}
	}
}

func (r *Room) broadcast(data []byte) {
	for client := range r.Clients {
		client.conn.WriteMessage(websocket.TextMessage, data)
	}
}
