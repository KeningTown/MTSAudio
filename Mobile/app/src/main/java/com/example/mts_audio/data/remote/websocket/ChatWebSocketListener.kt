package com.example.mts_audio.data.remote.websocket

import okhttp3.WebSocket

class ChatWebSocketListener(socketName: String, private val messageHandler: (String) -> Unit) : AppWebSocketListener(socketName) {
    override fun onMessage(webSocket: WebSocket, text: String) {
        super.onMessage(webSocket, text)
        messageHandler(text)
    }
}