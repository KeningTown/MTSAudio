package com.example.mts_audio.data.remote.websocket

import android.util.Log
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.WebSocket
import javax.inject.Inject

class WebSocketManager(
    private val okHttpClient: OkHttpClient,
){

    fun createWebSocket(baseUrl: String, endpoint: String): WebSocket {
        val url = "$baseUrl/$endpoint"
        Log.d("WEBSOCKET", url)
        val request = Request.Builder().url(url).build()
        return okHttpClient.newWebSocket(request, AppWebSocketListener(endpoint))
    }

}