package com.example.mts_audio.ui.model

class File(
    val audio_name : String,
    val chunk  : String,
    val done: Boolean
) {

    fun toJsonString(): String{
        return "{\"audio_name\": \"${audio_name}\", \"chunk\": \"${chunk}\",  \"done\": \"${done}\"}}"
    }

}