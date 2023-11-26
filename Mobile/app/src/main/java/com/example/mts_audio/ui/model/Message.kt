package com.example.mts_audio.ui.model

import com.example.mts_audio.data.local.User

class Message(
    val username: String,
    val msg: String,
) {
    fun toJsonString(): String{
        return "{\"username\": \"${username}\", \"msg\": \"${msg}\"}"
    }
}