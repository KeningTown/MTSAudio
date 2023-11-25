package com.example.mts_audio.data.remote.websocket

import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.WebSocket
import javax.inject.Inject

class WebSocketManager @Inject constructor(
    private val okHttpClient: OkHttpClient,
    private val baseUrl: String
){

    fun createWebSocket(endpoint: String): WebSocket {
        val url = "$baseUrl/$endpoint"
        val request = Request.Builder().url(url).build()
        return okHttpClient.newWebSocket(request, AppWebSocketListener(endpoint))
    }

}