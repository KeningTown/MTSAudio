package com.example.mts_audio.ui.model

class File(
    val audio_name : String,
    val chunk  : ByteArray,
) {

    fun toJsonString(): String{
        return "{\"audio_name\": \"${audio_name}\", \"chunk\": \"${chunk}\"}"
    }

}