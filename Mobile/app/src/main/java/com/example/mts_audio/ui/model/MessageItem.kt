package com.example.mts_audio.ui.model

import okhttp3.OkHttpClient

data class MessageItem(
    val message: Message,
    val isClient: Boolean,
)