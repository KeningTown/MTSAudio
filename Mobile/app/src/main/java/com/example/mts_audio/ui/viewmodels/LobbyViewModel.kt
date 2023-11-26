package com.example.mts_audio.ui.viewmodels

import android.media.MediaPlayer
import android.util.Log
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.example.mts_audio.data.model.Result
import com.example.mts_audio.data.remote.lobby.LobbyResult
import com.example.mts_audio.data.remote.websocket.ByteArrayDataSource
import com.example.mts_audio.data.remote.websocket.WebSocketManager
import com.example.mts_audio.data.repository.LobbyRepository
import com.example.mts_audio.data.repository.LocalUserRepository
import com.example.mts_audio.ui.model.File
import com.example.mts_audio.ui.model.Message
import com.example.mts_audio.ui.model.MessageItem
import com.google.gson.Gson
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import okhttp3.WebSocket
import javax.inject.Inject

@HiltViewModel
class LobbyViewModel @Inject constructor(
    private val lobbyRepository: LobbyRepository,
    private val localUserRepository: LocalUserRepository,
    private val webSocketManager: WebSocketManager
) : ViewModel() {

    private val _roomResult = MutableLiveData<LobbyResult>()
    val roomResult: LiveData<LobbyResult> = _roomResult

    private val _lobbyMessages = MutableLiveData<MessageItem>()
    val lobbyMessages: LiveData<MessageItem> = _lobbyMessages

    private var chatSocket: WebSocket? = null
    private var musicSocket: WebSocket? = null
    private var fileSocket: WebSocket? = null

    private lateinit var mediaPlayer: MediaPlayer
    private var music: ByteArray = byteArrayOf()

    private val userName = localUserRepository.getUserName()!!


    fun setRoom(roomId: String) {

        mediaPlayer = MediaPlayer()
        chatSocket = webSocketManager.createChatWebSocket("ws://10.0.2.2:80/ws", "${roomId}/chat") {
            onChatMessage(it)
        }
        chatSocket!!.send(accessTokenToJSON(localUserRepository.getAccessToken()!!))
        musicSocket = webSocketManager.createWebSocket("ws://10.0.2.2:80/ws", "${roomId}/track")
        musicSocket!!.send((accessTokenToJSON(localUserRepository.getAccessToken()!!)))
        fileSocket = webSocketManager.createChatWebSocket("ws://10.0.2.2:80/ws", "${roomId}/file"){
            onFileMessage(it)
        }
        fileSocket!!.send((accessTokenToJSON(localUserRepository.getAccessToken()!!)))
        fileSocket!!.send("{\"file_name\": \"видеообзор.mp3\"}")

    }

    private fun onChatMessage(message: String) {
        val outPut = parseJsonMessage(message)
        val isClient = (outPut.username == userName)
        GlobalScope.launch(Dispatchers.Main) {
            _lobbyMessages.value = MessageItem(Message(outPut.username, outPut.msg), isClient)
        }
        playAudio(music)
        Log.d("msg", "${outPut!!.toJsonString()}")
    }

    private fun onFileMessage(message: String) {
        val outPut = parseJsonFile(message)
        music.plus(outPut.chunk)
        Log.d("msg", "${outPut!!.toJsonString()}")
    }

    private fun accessTokenToJSON(accessToken: String): String {
        return "{\"access_token\": \"${accessToken}\"}"
    }

    fun parseJsonFile(jsonString: String): File {
        return try {
            val gson = Gson()
            gson.fromJson(jsonString, File::class.java)
        } catch (e: Exception) {
            File("", ByteArray(0))
        }
    }

    fun parseJsonMessage(jsonString: String): Message {
        return try {
            val gson = Gson()
            gson.fromJson(jsonString, Message::class.java)
        } catch (e: Exception) {
            Message(userName, "error")
        }
    }

    fun sendMessageToChat(message: String) {
        chatSocket?.send(Message(userName, message).toJsonString())
    }

    fun closeConnection() {
        chatSocket?.close(1000, "Close")
    }

    private fun playAudio(audioData: ByteArray) {
        try {
            mediaPlayer.stop();
            mediaPlayer.reset();

            mediaPlayer.setDataSource(ByteArrayDataSource(audioData));
            mediaPlayer.prepare();
            mediaPlayer.start();
        } catch (e: Exception) {
            Log.e("TAG", "Error playing audio $e");
        }
    }
}