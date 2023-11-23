package com.example.mts_audio.data.remote.auth

data class AuthResult(
    val success: AuthResponse? = null,
    val error: Int? = null
)
data class AuthResponse(
    val accessToken: String,
    val refreshToken: String,
    val user: User
)

data class User(
    val id: String,
    val email: String,
    val username: String,
    val password: String,
    val isActivated: Boolean,
    val activationLink: String
)