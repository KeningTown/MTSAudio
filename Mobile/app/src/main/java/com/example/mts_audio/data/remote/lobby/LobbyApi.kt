package com.example.mts_audio.data.remote.lobby

import com.example.mts_audio.data.remote.auth.AuthResponse
import com.example.mts_audio.data.remote.auth.UserRegistrationRequest
import kotlinx.coroutines.Deferred
import retrofit2.http.Body
import retrofit2.http.Headers
import retrofit2.http.POST

interface LobbyApi {

    @POST("/api/Room")
    @Headers("Content-Type: application/json")
    fun getRoom(): Deferred<LobbyResponse>

}