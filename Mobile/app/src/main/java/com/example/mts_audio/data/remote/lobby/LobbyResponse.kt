package com.example.mts_audio.data.remote.lobby

import com.example.mts_audio.data.remote.auth.AuthResponse

data class LobbyResult(
    val success: LobbyResponse? = null,
    val error: Int? = null
)
data class LobbyResponse(
    val roomId:String,
)
