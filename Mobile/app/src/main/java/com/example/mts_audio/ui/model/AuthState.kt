package com.example.mts_audio.ui.model

data class AuthFormState(
    val usernameError: Int? = null,
    val passwordError: Int? = null,
    val isDataValid: Boolean = false
)