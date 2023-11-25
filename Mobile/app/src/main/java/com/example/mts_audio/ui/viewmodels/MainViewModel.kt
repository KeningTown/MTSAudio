package com.example.mts_audio.ui.viewmodels

import android.util.Log
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.example.mts_audio.R
import com.example.mts_audio.data.model.Result
import com.example.mts_audio.data.remote.lobby.LobbyResult
import com.example.mts_audio.data.remote.websocket.WebSocketManager
import com.example.mts_audio.data.repository.AuthRepository
import com.example.mts_audio.data.repository.LobbyRepository
import com.example.mts_audio.data.repository.LocalUserRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import okhttp3.WebSocket
import okhttp3.WebSocketListener
import javax.inject.Inject

@HiltViewModel
class MainViewModel @Inject constructor(
    private val lobbyRepository: LobbyRepository,
    private val localUserRepository: LocalUserRepository,
    private val webSocketManager: WebSocketManager
) : ViewModel() {

    private val _roomResult = MutableLiveData<LobbyResult>()
    val roomResult: LiveData<LobbyResult> = _roomResult

    private var chatSocket:WebSocket? = null
    private var musicSocket:WebSocket? = null

    suspend fun getRoom() {
        val result = lobbyRepository.getRoom()

        if (result is Result.Success) {
            _roomResult.value = LobbyResult(success = result.data)
            chatSocket = webSocketManager.createWebSocket("ws://10.0.2.2:80/ws", "${result.data.roomId}/chat")
            chatSocket!!.send(accessTokenToJSON(localUserRepository.getAccessToken()!!))
            musicSocket = webSocketManager.createWebSocket("ws://10.0.2.2:80/ws", "${result.data.roomId}/music")
            musicSocket!!.send((accessTokenToJSON(localUserRepository.getAccessToken()!!)))
            Log.d("TAG", "lobby id ${result.data.roomId}")
        } else {
            _roomResult.value = LobbyResult(error = R.string.error_get_lobby)
        }
    }

    private fun accessTokenToJSON(accessToken:String): String {
       return "{\"access_token\": \"${accessToken}\"}"
    }

    fun closeConnection(){
        chatSocket?.close(1000, "Close")
    }
}