package com.example.mts_audio.data.remote.websocket

import okhttp3.WebSocket
import okhttp3.WebSocketListener
import okio.ByteString

class AppWebSocketListener(
    private val socketName: String
) : WebSocketListener() {

    companion object{
        val NORMAL_CLOSURE_STATUS = 1000
    }
    override fun onOpen(webSocket: WebSocket, response: okhttp3.Response) {
        super.onOpen(webSocket, response)
        println("$socketName opened")
    }

    override fun onMessage(webSocket: WebSocket, text: String) {
        super.onMessage(webSocket, text)
        println("$socketName received message: $text")
    }

    override fun onMessage(webSocket: WebSocket, bytes: ByteString) {
        super.onMessage(webSocket, bytes)
    }

    override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {
        super.onClosed(webSocket, NORMAL_CLOSURE_STATUS, reason)
        println("$socketName closed. Code: $code, Reason: $reason")
    }

    override fun onFailure(webSocket: WebSocket, t: Throwable, response: okhttp3.Response?) {
        super.onFailure(webSocket, t, response)
        println("$socketName failed. ${t.message}")
    }


}


