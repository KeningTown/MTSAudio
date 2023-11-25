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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024,
	WriteBufferSize: 1024 * 1024,
	//Solving cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wh WebsocketHandler) ChatConnect(roomId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer conn.Close()

		type accessToken struct {
			AccessToken string `json:"access_token"`
		}

		var token accessToken
		if err := conn.ReadJSON(&token); err != nil {
			log.Printf("failed to parse JSON: %s", err.Error())
			conn.WriteJSON(httputils.ResponseError{Error: "access_token required"})
			return
		}

		tokenData, err := tokens.ParseToken(token.AccessToken)
		if err != nil {
			conn.WriteJSON(httputils.ResponseError{Error: "invalid access token"})
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
			conn.WriteJSON(httputils.ResponseError{Error: "unexpected error"})
			log.Println("failed to send user amount")
			return
		}

		type message struct {
			Username string `json:"username"`
			Msg      string `json:"msg"`
		}

		for {
			var msg message

			if err := conn.ReadJSON(&msg); err != nil {
				log.Printf("chat websocket error: %s", err.Error())
				continue
			}
			msgData, err := json.Marshal(&msg)
			if err != nil {
				log.Printf("chat websocket error: %s", err.Error())
				continue
			}
			go room.broadcast(msgData)
		}
	}
}

func (wh WebsocketHandler) FileConnect(roomId, username, filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer conn.Close()

		client := Client{
			conn:     conn,
			username: username,
		}

		room := FileRooms[roomId]

		defer func() {
			room.Mu.Lock()
			delete(room.Clients, client)
			room.Mu.Unlock()
		}()

		room.Mu.Lock()

		room.Clients[client] = struct{}{}
		room.Mu.Unlock()

		type message struct {
			Filename string `json:"file_name"`
			OwnerId  uint   `json:"owner_id"`
		}

		for {
			var msg message
			if err := conn.ReadJSON(&msg); err != nil {
				continue
			}

			if msg.OwnerId != room.OwnerId {
				log.Println("no right to start sending file")
				continue
			}
			filePath := fmt.Sprintf("./static/%s", msg.Filename)
			absFilePath, err := filepath.Abs(filePath)

			fs, err := os.Open(absFilePath)

			if err != nil {
				conn.WriteJSON(httputils.ResponseError{
					Error: fmt.Sprintf("failed to open file: %s", absFilePath),
				})
				continue
			}
			defer fs.Close()

			type audioMessage struct {
				AudioName string `json:"audio_name"`
				Chunk     []byte `json:"chunk"`
			}

			for {
				buffer := make([]byte, 1024)
				n, err := fs.Read(buffer)
				if err != nil {
					if err == io.EOF {
						log.Printf("finished to send file: %s", filename)
						break
					}
					log.Print("fileWebsocket error: ", err.Error())
					break
				}

				msg := audioMessage{
					AudioName: filename,
					Chunk:     buffer[:n],
				}

				msgData, err := json.Marshal(msg)
				if err != nil {
					log.Printf("chat websocket error: %s", err.Error())
					continue
				}
				room.broadcast(msgData)
			}
		}
	}
}

func (r *Room) broadcast(data []byte) {
	for client := range r.Clients {
		client.conn.WriteMessage(websocket.TextMessage, data)
	}
}
