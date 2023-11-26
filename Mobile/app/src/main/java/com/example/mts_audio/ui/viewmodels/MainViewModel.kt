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


    suspend fun getRoom() {
        val result = lobbyRepository.getRoom()

        if (result is Result.Success) {
            _roomResult.value = LobbyResult(success = result.data)
            Log.d("TAG", "lobby id ${result.data.roomId}")
        } else {
            _roomResult.value = LobbyResult(error = R.string.error_get_lobby)
        }
    }


}