package com.example.mts_audio.data.repository

import android.util.Log
import com.example.mts_audio.data.model.Result
import com.example.mts_audio.data.remote.auth.AuthApi
import com.example.mts_audio.data.remote.auth.AuthResponse
import com.example.mts_audio.data.remote.auth.UserLoginRequest
import com.example.mts_audio.data.remote.lobby.LobbyApi
import com.example.mts_audio.data.remote.lobby.LobbyResponse
import kotlinx.coroutines.CoroutineDispatcher
import kotlinx.coroutines.withContext
import java.io.IOException

class LobbyRepository(
    private val lobbyDataSource: LobbyApi,
    private val ioDispatcher: CoroutineDispatcher,
) {

    suspend fun getRoom(): Result<LobbyResponse> {
        return try {
            Result.Success(
                withContext(ioDispatcher) {
                    val response = lobbyDataSource.getRoom()
                    response.await()
                }
            )
        } catch (e: Exception) {
            Log.d("TAG", e.message.toString())
            Result.Error(IOException("Error get lobby ", e))
        }
    }

}