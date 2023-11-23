package com.example.mts_audio.data.remote.auth

data class UserRegistrationRequest(
    val email: String,
    val password: String,
    val username: String
)

data class UserLoginRequest(
    val email: String,
    val password: String
)