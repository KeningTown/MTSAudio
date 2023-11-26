# Server
## Запуск
`docker-compose up`

<<<<<<< HEAD
Swagger: http://localhost:80/swagger/index.html
=======
Swagger: http://localhost:80/swagger/index.html

## Документация к Chat Websocket

Root: `ws://localhost:80/ws/:roomId/chat` - подключение к чату

Первое сообщение:
```json
{
    "access_token":"123213"
}
```
Response: `["user1", "user2"]`

Последующие сообщения:
```json
{
    "username" :"John Doe",
    "msg": "foo bar"
}
```

Response:

```json
{
    "username" :"John Doe",
    "msg": "foo bar"
}
```


## Документация к File Websocket
Root: `ws://localhost:80/ws/roomId/file`

После подключения отправить `access_token`:
```json
{
    "access_token": "token"
}
```
В ответ ничего не придет, но так подтвердится, что пользователь авторизован.
Владелец комнаты отсылает:
```json
{
    "file_name": "track.mp3"
}
```
Если будет указано id **не владельца** комнаты, то ничего не произойдет.
Response:
```json
{
    "audio_name": "звуки фонка.mp3",
    "chunk": "*байты файла*",
    "done":false //когда заканчивается загрузка, done = true
}
```

## Документация по Track Websocket
Root: `ws://localhost:80/ws/roomId/file`

После подключения отправить `access_token`:
```json
{
    "access_token": "token"
}
```
В ответ ничего не придет, но так подтвердится, что пользователь авторизован.
Владелец комнаты отсылает:
```json
{
    "play_misic": true
}
```
Response:
```json
{
    "play_misic": true
}
```
>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
